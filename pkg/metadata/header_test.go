package metadata_test

import (
	"encoding/hex"
	"math"
	"reflect"
	"testing"
	"time"

	gogotypes "github.com/gogo/protobuf/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/merkle"
	"github.com/tendermint/tendermint/crypto/tmhash"
	test "github.com/tendermint/tendermint/internal/test/factory"
	"github.com/tendermint/tendermint/libs/bytes"
	"github.com/tendermint/tendermint/pkg/metadata"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmversion "github.com/tendermint/tendermint/proto/tendermint/version"
	"github.com/tendermint/tendermint/version"
)

var nilBytes []byte

func TestNilHeaderHashDoesntCrash(t *testing.T) {
	assert.Equal(t, nilBytes, []byte((*metadata.Header)(nil).Hash()))
	assert.Equal(t, nilBytes, []byte((new(metadata.Header)).Hash()))
}

func TestHeaderHash(t *testing.T) {
	testCases := []struct {
		desc       string
		header     *metadata.Header
		expectHash bytes.HexBytes
	}{
		{"Generates expected hash", &metadata.Header{
			Version:            version.Consensus{Block: 1, App: 2},
			ChainID:            "chainId",
			Height:             3,
			Time:               time.Date(2019, 10, 13, 16, 14, 44, 0, time.UTC),
			LastBlockID:        test.MakeBlockIDWithHash(make([]byte, tmhash.Size)),
			LastCommitHash:     tmhash.Sum([]byte("last_commit_hash")),
			DataHash:           tmhash.Sum([]byte("data_hash")),
			ValidatorsHash:     tmhash.Sum([]byte("validators_hash")),
			NextValidatorsHash: tmhash.Sum([]byte("next_validators_hash")),
			ConsensusHash:      tmhash.Sum([]byte("consensus_hash")),
			AppHash:            tmhash.Sum([]byte("app_hash")),
			LastResultsHash:    tmhash.Sum([]byte("last_results_hash")),
			EvidenceHash:       tmhash.Sum([]byte("evidence_hash")),
			ProposerAddress:    crypto.AddressHash([]byte("proposer_address")),
		}, hexBytesFromString("B8C1FA74E943A05664AD19C97D6D89EED19400D6749D912C2F3A4AA15B3D8E92")},
		{"nil header yields nil", nil, nil},
		{"nil ValidatorsHash yields nil", &metadata.Header{
			Version:            version.Consensus{Block: 1, App: 2},
			ChainID:            "chainId",
			Height:             3,
			Time:               time.Date(2019, 10, 13, 16, 14, 44, 0, time.UTC),
			LastBlockID:        test.MakeBlockIDWithHash(make([]byte, tmhash.Size)),
			LastCommitHash:     tmhash.Sum([]byte("last_commit_hash")),
			DataHash:           tmhash.Sum([]byte("data_hash")),
			ValidatorsHash:     nil,
			NextValidatorsHash: tmhash.Sum([]byte("next_validators_hash")),
			ConsensusHash:      tmhash.Sum([]byte("consensus_hash")),
			AppHash:            tmhash.Sum([]byte("app_hash")),
			LastResultsHash:    tmhash.Sum([]byte("last_results_hash")),
			EvidenceHash:       tmhash.Sum([]byte("evidence_hash")),
			ProposerAddress:    crypto.AddressHash([]byte("proposer_address")),
		}, nil},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			assert.Equal(t, tc.expectHash, tc.header.Hash())

			// We also make sure that all fields are hashed in struct order, and that all
			// fields in the test struct are non-zero.
			if tc.header != nil && tc.expectHash != nil {
				byteSlices := [][]byte{}

				s := reflect.ValueOf(*tc.header)
				for i := 0; i < s.NumField(); i++ {
					f := s.Field(i)

					assert.False(t, f.IsZero(), "Found zero-valued field %v",
						s.Type().Field(i).Name)

					switch f := f.Interface().(type) {
					case int64, bytes.HexBytes, string:
						byteSlices = append(byteSlices, metadata.CdcEncode(f))
					case time.Time:
						bz, err := gogotypes.StdTimeMarshal(f)
						require.NoError(t, err)
						byteSlices = append(byteSlices, bz)
					case version.Consensus:
						pbc := tmversion.Consensus{
							Block: f.Block,
							App:   f.App,
						}
						bz, err := pbc.Marshal()
						require.NoError(t, err)
						byteSlices = append(byteSlices, bz)
					case metadata.BlockID:
						pbbi := f.ToProto()
						bz, err := pbbi.Marshal()
						require.NoError(t, err)
						byteSlices = append(byteSlices, bz)
					default:
						t.Errorf("unknown type %T", f)
					}
				}
				assert.Equal(t,
					bytes.HexBytes(merkle.HashFromByteSlices(byteSlices)), tc.header.Hash())
			}
		})
	}
}

func TestMaxHeaderBytes(t *testing.T) {
	// Construct a UTF-8 string of MaxChainIDLen length using the supplementary
	// characters.
	// Each supplementary character takes 4 bytes.
	// http://www.i18nguy.com/unicode/supplementary-test.html
	maxChainID := ""
	for i := 0; i < metadata.MaxChainIDLen; i++ {
		maxChainID += "𠜎"
	}

	// time is varint encoded so need to pick the max.
	// year int, month Month, day, hour, min, sec, nsec int, loc *Location
	timestamp := time.Date(math.MaxInt64, 0, 0, 0, 0, 0, math.MaxInt64, time.UTC)

	h := metadata.Header{
		Version: version.Consensus{Block: math.MaxInt64, App: math.MaxInt64},
		ChainID: maxChainID,
		Height:  math.MaxInt64,
		Time:    timestamp,
		LastBlockID: metadata.BlockID{
			Hash:          make([]byte, tmhash.Size),
			PartSetHeader: metadata.PartSetHeader{math.MaxInt32, make([]byte, tmhash.Size)},
		},
		LastCommitHash:     tmhash.Sum([]byte("last_commit_hash")),
		DataHash:           tmhash.Sum([]byte("data_hash")),
		ValidatorsHash:     tmhash.Sum([]byte("validators_hash")),
		NextValidatorsHash: tmhash.Sum([]byte("next_validators_hash")),
		ConsensusHash:      tmhash.Sum([]byte("consensus_hash")),
		AppHash:            tmhash.Sum([]byte("app_hash")),
		LastResultsHash:    tmhash.Sum([]byte("last_results_hash")),
		EvidenceHash:       tmhash.Sum([]byte("evidence_hash")),
		ProposerAddress:    crypto.AddressHash([]byte("proposer_address")),
	}

	bz, err := h.ToProto().Marshal()
	require.NoError(t, err)

	assert.EqualValues(t, metadata.MaxHeaderBytes, int64(len(bz)))
}

func randCommit(now time.Time) *metadata.Commit {
	lastID := test.MakeBlockID()
	h := int64(3)
	voteSet, _, vals := test.RandVoteSet(h-1, 1, tmproto.PrecommitType, 10, 1)
	commit, err := test.MakeCommit(lastID, h-1, 1, voteSet, vals, now)
	if err != nil {
		panic(err)
	}
	return commit
}

func hexBytesFromString(s string) bytes.HexBytes {
	b, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return bytes.HexBytes(b)
}

func TestHeader_ValidateBasic(t *testing.T) {
	testCases := []struct {
		name      string
		header    metadata.Header
		expectErr bool
		errString string
	}{
		{
			"invalid version block",
			metadata.Header{Version: version.Consensus{Block: version.BlockProtocol + 1}},
			true, "block protocol is incorrect",
		},
		{
			"invalid chain ID length",
			metadata.Header{
				Version: version.Consensus{Block: version.BlockProtocol},
				ChainID: string(make([]byte, metadata.MaxChainIDLen+1)),
			},
			true, "chainID is too long",
		},
		{
			"invalid height (negative)",
			metadata.Header{
				Version: version.Consensus{Block: version.BlockProtocol},
				ChainID: string(make([]byte, metadata.MaxChainIDLen)),
				Height:  -1,
			},
			true, "negative Height",
		},
		{
			"invalid height (zero)",
			metadata.Header{
				Version: version.Consensus{Block: version.BlockProtocol},
				ChainID: string(make([]byte, metadata.MaxChainIDLen)),
				Height:  0,
			},
			true, "zero Height",
		},
		{
			"invalid block ID hash",
			metadata.Header{
				Version: version.Consensus{Block: version.BlockProtocol},
				ChainID: string(make([]byte, metadata.MaxChainIDLen)),
				Height:  1,
				LastBlockID: metadata.BlockID{
					Hash: make([]byte, tmhash.Size+1),
				},
			},
			true, "wrong Hash",
		},
		{
			"invalid block ID parts header hash",
			metadata.Header{
				Version: version.Consensus{Block: version.BlockProtocol},
				ChainID: string(make([]byte, metadata.MaxChainIDLen)),
				Height:  1,
				LastBlockID: metadata.BlockID{
					Hash: make([]byte, tmhash.Size),
					PartSetHeader: metadata.PartSetHeader{
						Hash: make([]byte, tmhash.Size+1),
					},
				},
			},
			true, "wrong PartSetHeader",
		},
		{
			"invalid last commit hash",
			metadata.Header{
				Version: version.Consensus{Block: version.BlockProtocol},
				ChainID: string(make([]byte, metadata.MaxChainIDLen)),
				Height:  1,
				LastBlockID: metadata.BlockID{
					Hash: make([]byte, tmhash.Size),
					PartSetHeader: metadata.PartSetHeader{
						Hash: make([]byte, tmhash.Size),
					},
				},
				LastCommitHash: make([]byte, tmhash.Size+1),
			},
			true, "wrong LastCommitHash",
		},
		{
			"invalid data hash",
			metadata.Header{
				Version: version.Consensus{Block: version.BlockProtocol},
				ChainID: string(make([]byte, metadata.MaxChainIDLen)),
				Height:  1,
				LastBlockID: metadata.BlockID{
					Hash: make([]byte, tmhash.Size),
					PartSetHeader: metadata.PartSetHeader{
						Hash: make([]byte, tmhash.Size),
					},
				},
				LastCommitHash: make([]byte, tmhash.Size),
				DataHash:       make([]byte, tmhash.Size+1),
			},
			true, "wrong DataHash",
		},
		{
			"invalid evidence hash",
			metadata.Header{
				Version: version.Consensus{Block: version.BlockProtocol},
				ChainID: string(make([]byte, metadata.MaxChainIDLen)),
				Height:  1,
				LastBlockID: metadata.BlockID{
					Hash: make([]byte, tmhash.Size),
					PartSetHeader: metadata.PartSetHeader{
						Hash: make([]byte, tmhash.Size),
					},
				},
				LastCommitHash: make([]byte, tmhash.Size),
				DataHash:       make([]byte, tmhash.Size),
				EvidenceHash:   make([]byte, tmhash.Size+1),
			},
			true, "wrong EvidenceHash",
		},
		{
			"invalid proposer address",
			metadata.Header{
				Version: version.Consensus{Block: version.BlockProtocol},
				ChainID: string(make([]byte, metadata.MaxChainIDLen)),
				Height:  1,
				LastBlockID: metadata.BlockID{
					Hash: make([]byte, tmhash.Size),
					PartSetHeader: metadata.PartSetHeader{
						Hash: make([]byte, tmhash.Size),
					},
				},
				LastCommitHash:  make([]byte, tmhash.Size),
				DataHash:        make([]byte, tmhash.Size),
				EvidenceHash:    make([]byte, tmhash.Size),
				ProposerAddress: make([]byte, crypto.AddressSize+1),
			},
			true, "invalid ProposerAddress length",
		},
		{
			"invalid validator hash",
			metadata.Header{
				Version: version.Consensus{Block: version.BlockProtocol},
				ChainID: string(make([]byte, metadata.MaxChainIDLen)),
				Height:  1,
				LastBlockID: metadata.BlockID{
					Hash: make([]byte, tmhash.Size),
					PartSetHeader: metadata.PartSetHeader{
						Hash: make([]byte, tmhash.Size),
					},
				},
				LastCommitHash:  make([]byte, tmhash.Size),
				DataHash:        make([]byte, tmhash.Size),
				EvidenceHash:    make([]byte, tmhash.Size),
				ProposerAddress: make([]byte, crypto.AddressSize),
				ValidatorsHash:  make([]byte, tmhash.Size+1),
			},
			true, "wrong ValidatorsHash",
		},
		{
			"invalid next validator hash",
			metadata.Header{
				Version: version.Consensus{Block: version.BlockProtocol},
				ChainID: string(make([]byte, metadata.MaxChainIDLen)),
				Height:  1,
				LastBlockID: metadata.BlockID{
					Hash: make([]byte, tmhash.Size),
					PartSetHeader: metadata.PartSetHeader{
						Hash: make([]byte, tmhash.Size),
					},
				},
				LastCommitHash:     make([]byte, tmhash.Size),
				DataHash:           make([]byte, tmhash.Size),
				EvidenceHash:       make([]byte, tmhash.Size),
				ProposerAddress:    make([]byte, crypto.AddressSize),
				ValidatorsHash:     make([]byte, tmhash.Size),
				NextValidatorsHash: make([]byte, tmhash.Size+1),
			},
			true, "wrong NextValidatorsHash",
		},
		{
			"invalid consensus hash",
			metadata.Header{
				Version: version.Consensus{Block: version.BlockProtocol},
				ChainID: string(make([]byte, metadata.MaxChainIDLen)),
				Height:  1,
				LastBlockID: metadata.BlockID{
					Hash: make([]byte, tmhash.Size),
					PartSetHeader: metadata.PartSetHeader{
						Hash: make([]byte, tmhash.Size),
					},
				},
				LastCommitHash:     make([]byte, tmhash.Size),
				DataHash:           make([]byte, tmhash.Size),
				EvidenceHash:       make([]byte, tmhash.Size),
				ProposerAddress:    make([]byte, crypto.AddressSize),
				ValidatorsHash:     make([]byte, tmhash.Size),
				NextValidatorsHash: make([]byte, tmhash.Size),
				ConsensusHash:      make([]byte, tmhash.Size+1),
			},
			true, "wrong ConsensusHash",
		},
		{
			"invalid last results hash",
			metadata.Header{
				Version: version.Consensus{Block: version.BlockProtocol},
				ChainID: string(make([]byte, metadata.MaxChainIDLen)),
				Height:  1,
				LastBlockID: metadata.BlockID{
					Hash: make([]byte, tmhash.Size),
					PartSetHeader: metadata.PartSetHeader{
						Hash: make([]byte, tmhash.Size),
					},
				},
				LastCommitHash:     make([]byte, tmhash.Size),
				DataHash:           make([]byte, tmhash.Size),
				EvidenceHash:       make([]byte, tmhash.Size),
				ProposerAddress:    make([]byte, crypto.AddressSize),
				ValidatorsHash:     make([]byte, tmhash.Size),
				NextValidatorsHash: make([]byte, tmhash.Size),
				ConsensusHash:      make([]byte, tmhash.Size),
				LastResultsHash:    make([]byte, tmhash.Size+1),
			},
			true, "wrong LastResultsHash",
		},
		{
			"valid header",
			metadata.Header{
				Version: version.Consensus{Block: version.BlockProtocol},
				ChainID: string(make([]byte, metadata.MaxChainIDLen)),
				Height:  1,
				LastBlockID: metadata.BlockID{
					Hash: make([]byte, tmhash.Size),
					PartSetHeader: metadata.PartSetHeader{
						Hash: make([]byte, tmhash.Size),
					},
				},
				LastCommitHash:     make([]byte, tmhash.Size),
				DataHash:           make([]byte, tmhash.Size),
				EvidenceHash:       make([]byte, tmhash.Size),
				ProposerAddress:    make([]byte, crypto.AddressSize),
				ValidatorsHash:     make([]byte, tmhash.Size),
				NextValidatorsHash: make([]byte, tmhash.Size),
				ConsensusHash:      make([]byte, tmhash.Size),
				LastResultsHash:    make([]byte, tmhash.Size),
			},
			false, "",
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			err := tc.header.ValidateBasic()
			if tc.expectErr {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.errString)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestHeaderProto(t *testing.T) {
	h1 := test.MakeRandomHeader()
	tc := []struct {
		msg     string
		h1      *metadata.Header
		expPass bool
	}{
		{"success", h1, true},
		{"failure empty Header", &metadata.Header{}, false},
	}

	for _, tt := range tc {
		tt := tt
		t.Run(tt.msg, func(t *testing.T) {
			pb := tt.h1.ToProto()
			h, err := metadata.HeaderFromProto(pb)
			if tt.expPass {
				require.NoError(t, err, tt.msg)
				require.Equal(t, tt.h1, &h, tt.msg)
			} else {
				require.Error(t, err, tt.msg)
			}

		})
	}
}

func TestHeaderHashVector(t *testing.T) {
	chainID := "test"
	h := metadata.Header{
		Version:            version.Consensus{Block: 1, App: 1},
		ChainID:            chainID,
		Height:             50,
		Time:               time.Date(math.MaxInt64, 0, 0, 0, 0, 0, math.MaxInt64, time.UTC),
		LastBlockID:        metadata.BlockID{},
		LastCommitHash:     []byte("f2564c78071e26643ae9b3e2a19fa0dc10d4d9e873aa0be808660123f11a1e78"),
		DataHash:           []byte("f2564c78071e26643ae9b3e2a19fa0dc10d4d9e873aa0be808660123f11a1e78"),
		ValidatorsHash:     []byte("f2564c78071e26643ae9b3e2a19fa0dc10d4d9e873aa0be808660123f11a1e78"),
		NextValidatorsHash: []byte("f2564c78071e26643ae9b3e2a19fa0dc10d4d9e873aa0be808660123f11a1e78"),
		ConsensusHash:      []byte("f2564c78071e26643ae9b3e2a19fa0dc10d4d9e873aa0be808660123f11a1e78"),
		AppHash:            []byte("f2564c78071e26643ae9b3e2a19fa0dc10d4d9e873aa0be808660123f11a1e78"),

		LastResultsHash: []byte("f2564c78071e26643ae9b3e2a19fa0dc10d4d9e873aa0be808660123f11a1e78"),

		EvidenceHash:    []byte("f2564c78071e26643ae9b3e2a19fa0dc10d4d9e873aa0be808660123f11a1e78"),
		ProposerAddress: []byte("2915b7b15f979e48ebc61774bb1d86ba3136b7eb"),
	}

	testCases := []struct {
		header   metadata.Header
		expBytes string
	}{
		{header: h, expBytes: "87b6117ac7f827d656f178a3d6d30b24b205db2b6a3a053bae8baf4618570bfc"},
	}

	for _, tc := range testCases {
		hash := tc.header.Hash()
		require.Equal(t, tc.expBytes, hex.EncodeToString(hash))
	}
}

func TestSignedHeaderValidateBasic(t *testing.T) {
	commit := test.MakeRandomCommit(time.Now())
	chainID := "𠜎"
	timestamp := time.Date(math.MaxInt64, 0, 0, 0, 0, 0, math.MaxInt64, time.UTC)
	h := metadata.Header{
		Version:            version.Consensus{Block: version.BlockProtocol, App: math.MaxInt64},
		ChainID:            chainID,
		Height:             commit.Height,
		Time:               timestamp,
		LastBlockID:        commit.BlockID,
		LastCommitHash:     commit.Hash(),
		DataHash:           commit.Hash(),
		ValidatorsHash:     commit.Hash(),
		NextValidatorsHash: commit.Hash(),
		ConsensusHash:      commit.Hash(),
		AppHash:            commit.Hash(),
		LastResultsHash:    commit.Hash(),
		EvidenceHash:       commit.Hash(),
		ProposerAddress:    crypto.AddressHash([]byte("proposer_address")),
	}

	validSignedHeader := metadata.SignedHeader{Header: &h, Commit: commit}
	validSignedHeader.Commit.BlockID.Hash = validSignedHeader.Hash()
	invalidSignedHeader := metadata.SignedHeader{}

	testCases := []struct {
		testName  string
		shHeader  *metadata.Header
		shCommit  *metadata.Commit
		expectErr bool
	}{
		{"Valid Signed Header", validSignedHeader.Header, validSignedHeader.Commit, false},
		{"Invalid Signed Header", invalidSignedHeader.Header, validSignedHeader.Commit, true},
		{"Invalid Signed Header", validSignedHeader.Header, invalidSignedHeader.Commit, true},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.testName, func(t *testing.T) {
			sh := metadata.SignedHeader{
				Header: tc.shHeader,
				Commit: tc.shCommit,
			}
			err := sh.ValidateBasic(validSignedHeader.Header.ChainID)
			assert.Equalf(
				t,
				tc.expectErr,
				err != nil,
				"Validate Basic had an unexpected result",
				err,
			)
		})
	}
}

func TestSignedHeaderProtoBuf(t *testing.T) {
	commit := test.MakeRandomCommit(time.Now())
	h := test.MakeRandomHeader()

	sh := metadata.SignedHeader{Header: h, Commit: commit}

	testCases := []struct {
		msg     string
		sh1     *metadata.SignedHeader
		expPass bool
	}{
		{"empty SignedHeader 2", &metadata.SignedHeader{}, true},
		{"success", &sh, true},
		{"failure nil", nil, false},
	}
	for _, tc := range testCases {
		protoSignedHeader := tc.sh1.ToProto()

		sh, err := metadata.SignedHeaderFromProto(protoSignedHeader)

		if tc.expPass {
			require.NoError(t, err, tc.msg)
			require.Equal(t, tc.sh1, sh, tc.msg)
		} else {
			require.Error(t, err, tc.msg)
		}
	}
}