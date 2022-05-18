(window.webpackJsonp=window.webpackJsonp||[]).push([[193],{767:function(e,t,i){"use strict";i.r(t);var n=i(1),a=Object(n.a)({},(function(){var e=this,t=e.$createElement,i=e._self._c||t;return i("ContentSlotsDistributor",{attrs:{"slot-key":e.$parent.slotKey}},[i("h1",{attrs:{id:"evidence"}},[i("a",{staticClass:"header-anchor",attrs:{href:"#evidence"}},[e._v("#")]),e._v(" Evidence")]),e._v(" "),i("p",[e._v('Evidence is an important component of Tendermint\'s security model. Whilst the core\nconsensus protocol provides correctness guarantees for state machine replication\nthat can tolerate less than 1/3 failures, the evidence system looks to detect and\ngossip byzantine faults whose combined power is greater than  or equal to 1/3. It is worth noting that\nthe evidence system is designed purely to detect possible attacks, gossip them,\ncommit them on chain and inform the application running on top of Tendermint.\nEvidence in itself does not punish "bad actors", this is left to the discretion\nof the application. A common form of punishment is slashing where the validators\nthat were caught violating the protocol have all or a portion of their voting\npower removed. Evidence, given the assumption that 1/3+ of the network is still\nbyzantine, is susceptible to censorship and should therefore be considered added\nsecurity on a "best effort" basis.')]),e._v(" "),i("p",[e._v("This document walks through the various forms of evidence, how they are detected,\ngossiped, verified and committed.")]),e._v(" "),i("blockquote",[i("p",[e._v("NOTE: Evidence here is internal to tendermint and should not be confused with\napplication evidence")])]),e._v(" "),i("h2",{attrs:{id:"detection"}},[i("a",{staticClass:"header-anchor",attrs:{href:"#detection"}},[e._v("#")]),e._v(" Detection")]),e._v(" "),i("h3",{attrs:{id:"equivocation"}},[i("a",{staticClass:"header-anchor",attrs:{href:"#equivocation"}},[e._v("#")]),e._v(" Equivocation")]),e._v(" "),i("p",[e._v("Equivocation is the most fundamental of byzantine faults. Simply put, to prevent\nreplication of state across all nodes, a validator tries to convince some subset\nof nodes to commit one block whilst convincing another subset to commit a\ndifferent block. This is achieved by double voting (hence\n"),i("code",[e._v("DuplicateVoteEvidence")]),e._v("). A successful duplicate vote attack requires greater\nthan 1/3 voting power and a (temporary) network partition between the aforementioned\nsubsets. This is because in consensus, votes are gossiped around. When a node\nobserves two conflicting votes from the same peer, it will use the two votes of\nevidence and begin gossiping this evidence to other nodes. "),i("a",{attrs:{href:"#duplicatevoteevidence"}},[e._v("Verification")]),e._v(" is addressed further down.")]),e._v(" "),i("tm-code-block",{staticClass:"codeblock",attrs:{language:"go",base64:"dHlwZSBEdXBsaWNhdGVWb3RlRXZpZGVuY2Ugc3RydWN0IHsKICAgIFZvdGVBIFZvdGUKICAgIFZvdGVCIFZvdGUKCiAgICAvLyBhbmQgYWJjaSBzcGVjaWZpYyBmaWVsZHMKfQo="}}),e._v(" "),i("h3",{attrs:{id:"light-client-attacks"}},[i("a",{staticClass:"header-anchor",attrs:{href:"#light-client-attacks"}},[e._v("#")]),e._v(" Light Client Attacks")]),e._v(" "),i("p",[e._v("Light clients also comply with the 1/3+ security model, however, by using a\ndifferent, more lightweight verification method they are subject to a\ndifferent kind of 1/3+ attack whereby the byzantine validators could sign an\nalternative light block that the light client will think is valid. Detection,\nexplained in greater detail\n"),i("RouterLink",{attrs:{to:"/spec/light-client/detection/detection_003_reviewed.html"}},[e._v("here")]),e._v(', involves comparison\nwith multiple other nodes in the hope that at least one is "honest". An "honest"\nnode will return a challenging light block for the light client to validate. If\nthis challenging light block also meets the\n'),i("RouterLink",{attrs:{to:"/spec/light-client/verification/verification_001_published.html"}},[e._v("validation criteria")]),e._v('\nthen the light client sends the "forged" light block to the node.\n'),i("a",{attrs:{href:"#lightclientattackevidence"}},[e._v("Verification")]),e._v(" is addressed further down.")],1),e._v(" "),i("tm-code-block",{staticClass:"codeblock",attrs:{language:"go",base64:"dHlwZSBMaWdodENsaWVudEF0dGFja0V2aWRlbmNlIHN0cnVjdCB7CiAgICBDb25mbGljdGluZ0Jsb2NrIExpZ2h0QmxvY2sKICAgIENvbW1vbkhlaWdodCBpbnQ2NAoKICAgICAgLy8gYW5kIGFiY2kgc3BlY2lmaWMgZmllbGRzCn0K"}}),e._v(" "),i("h2",{attrs:{id:"verification"}},[i("a",{staticClass:"header-anchor",attrs:{href:"#verification"}},[e._v("#")]),e._v(" Verification")]),e._v(" "),i("p",[e._v("If a node receives evidence, it will first try to verify it, then persist it.\nEvidence of byzantine behavior should only be committed once (uniqueness) and\nshould be committed within a certain period from the point that it occurred\n(timely). Timelines is defined by the "),i("code",[e._v("EvidenceParams")]),e._v(": "),i("code",[e._v("MaxAgeNumBlocks")]),e._v(" and\n"),i("code",[e._v("MaxAgeDuration")]),e._v(". In Proof of Stake chains where validators are bonded, evidence\nage should be less than the unbonding period so validators still can be\npunished. Given these two propoerties the following initial checks are made.")]),e._v(" "),i("ol",[i("li",[i("p",[e._v("Has the evidence expired? This is done by taking the height of the "),i("code",[e._v("Vote")]),e._v("\nwithin "),i("code",[e._v("DuplicateVoteEvidence")]),e._v(" or "),i("code",[e._v("CommonHeight")]),e._v(" within\n"),i("code",[e._v("LightClientAttakEvidence")]),e._v(". The evidence height is then used to retrieve the\nheader and thus the time of the block that corresponds to the evidence. If\n"),i("code",[e._v("CurrentHeight - MaxAgeNumBlocks > EvidenceHeight")]),e._v(" && "),i("code",[e._v("CurrentTime - MaxAgeDuration > EvidenceTime")]),e._v(", the evidence is considered expired and\nignored.")])]),e._v(" "),i("li",[i("p",[e._v("Has the evidence already been committed? The evidence pool tracks the hash of\nall committed evidence and uses this to determine uniqueness. If a new\nevidence has the same hash as a committed one, the new evidence will be\nignored.")])])]),e._v(" "),i("h3",{attrs:{id:"duplicatevoteevidence"}},[i("a",{staticClass:"header-anchor",attrs:{href:"#duplicatevoteevidence"}},[e._v("#")]),e._v(" DuplicateVoteEvidence")]),e._v(" "),i("p",[e._v("Valid "),i("code",[e._v("DuplicateVoteEvidence")]),e._v(" must adhere to the following rules:")]),e._v(" "),i("ul",[i("li",[i("p",[e._v("Validator Address, Height, Round and Type must be the same for both votes")])]),e._v(" "),i("li",[i("p",[e._v("BlockID must be different for both votes (BlockID can be for a nil block)")])]),e._v(" "),i("li",[i("p",[e._v("Validator must have been in the validator set at that height")])]),e._v(" "),i("li",[i("p",[e._v("Vote signature must be correctly signed. This also uses "),i("code",[e._v("ChainID")]),e._v(" so we know\nthat the fault occurred on this chain")])])]),e._v(" "),i("h3",{attrs:{id:"lightclientattackevidence"}},[i("a",{staticClass:"header-anchor",attrs:{href:"#lightclientattackevidence"}},[e._v("#")]),e._v(" LightClientAttackEvidence")]),e._v(" "),i("p",[e._v("Valid Light Client Attack Evidence must adhere to the following rules:")]),e._v(" "),i("ul",[i("li",[i("p",[e._v("If the header of the light block is invalid, thus indicating a lunatic attack,\nthe node must check that they can use "),i("code",[e._v("verifySkipping")]),e._v(" from their header at\nthe common height to the conflicting header")])]),e._v(" "),i("li",[i("p",[e._v("If the header is valid, then the validator sets are the same and this is\neither a form of equivocation or amnesia. We therefore check that 2/3 of the\nvalidator set also signed the conflicting header.")])]),e._v(" "),i("li",[i("p",[e._v("The nodes own header at the same height as the conflicting header must have a\ndifferent hash to the conflicting header.")])]),e._v(" "),i("li",[i("p",[e._v("If the nodes latest header is less in height to the conflicting header, then\nthe node must check that the conflicting block has a time that is less than\nthis latest header (This is a forward lunatic attack).")])])]),e._v(" "),i("h2",{attrs:{id:"gossiping"}},[i("a",{staticClass:"header-anchor",attrs:{href:"#gossiping"}},[e._v("#")]),e._v(" Gossiping")]),e._v(" "),i("p",[e._v("If a node verifies evidence it then broadcasts it to all peers, continously sending\nthe same evidence once every 10 seconds until the evidence is seen on chain or\nexpires.")]),e._v(" "),i("h2",{attrs:{id:"commiting-on-chain"}},[i("a",{staticClass:"header-anchor",attrs:{href:"#commiting-on-chain"}},[e._v("#")]),e._v(" Commiting on Chain")]),e._v(" "),i("p",[e._v("Evidence takes strict priority over regular transactions, thus a block is filled\nwith evidence first and transactions take up the remainder of the space. To\nmitigate the threat of an already punished node from spamming the network with\nmore evidence, the size of the evidence in a block can be capped by\n"),i("code",[e._v("EvidenceParams.MaxBytes")]),e._v(". Nodes receiving blocks with evidence will validate\nthe evidence before sending "),i("code",[e._v("Prevote")]),e._v(" and "),i("code",[e._v("Precommit")]),e._v(" votes. The evidence pool\nwill usually cache verifications so that this process is much quicker.")]),e._v(" "),i("h2",{attrs:{id:"sending-evidence-to-the-application"}},[i("a",{staticClass:"header-anchor",attrs:{href:"#sending-evidence-to-the-application"}},[e._v("#")]),e._v(" Sending Evidence to the Application")]),e._v(" "),i("p",[e._v("After evidence is committed, the block is then processed by the block executor\nwhich delivers the evidence to the application via "),i("code",[e._v("EndBlock")]),e._v(". Evidence is\nstripped of the actual proof, split up per faulty validator and only the\nvalidator, height, time and evidence type is sent.")]),e._v(" "),i("tm-code-block",{staticClass:"codeblock",attrs:{language:"proto",base64:"ZW51bSBFdmlkZW5jZVR5cGUgewogIFVOS05PV04gICAgICAgICAgICAgPSAwOwogIERVUExJQ0FURV9WT1RFICAgICAgPSAxOwogIExJR0hUX0NMSUVOVF9BVFRBQ0sgPSAyOwp9CgptZXNzYWdlIEV2aWRlbmNlIHsKICBFdmlkZW5jZVR5cGUgdHlwZSA9IDE7CiAgLy8gVGhlIG9mZmVuZGluZyB2YWxpZGF0b3IKICBWYWxpZGF0b3IgdmFsaWRhdG9yID0gMiBbKGdvZ29wcm90by5udWxsYWJsZSkgPSBmYWxzZV07CiAgLy8gVGhlIGhlaWdodCB3aGVuIHRoZSBvZmZlbnNlIG9jY3VycmVkCiAgaW50NjQgaGVpZ2h0ID0gMzsKICAvLyBUaGUgY29ycmVzcG9uZGluZyB0aW1lIHdoZXJlIHRoZSBvZmZlbnNlIG9jY3VycmVkCiAgZ29vZ2xlLnByb3RvYnVmLlRpbWVzdGFtcCB0aW1lID0gNCBbCiAgICAoZ29nb3Byb3RvLm51bGxhYmxlKSA9IGZhbHNlLCAoZ29nb3Byb3RvLnN0ZHRpbWUpID0gdHJ1ZV07CiAgLy8gVG90YWwgdm90aW5nIHBvd2VyIG9mIHRoZSB2YWxpZGF0b3Igc2V0IGluIGNhc2UgdGhlIEFCQ0kgYXBwbGljYXRpb24gZG9lcwogIC8vIG5vdCBzdG9yZSBoaXN0b3JpY2FsIHZhbGlkYXRvcnMuCiAgLy8gaHR0cHM6Ly9naXRodWIuY29tL3RlbmRlcm1pbnQvdGVuZGVybWludC9pc3N1ZXMvNDU4MQogIGludDY0IHRvdGFsX3ZvdGluZ19wb3dlciA9IDU7Cn0K"}}),e._v(" "),i("p",[i("code",[e._v("DuplicateVoteEvidence")]),e._v(" and "),i("code",[e._v("LightClientAttackEvidence")]),e._v(" are self-contained in\nthe sense that the evidence can be used to derive the "),i("code",[e._v("abci.Evidence")]),e._v(" that is\nsent to the application. Because of this, extra fields are necessary:")]),e._v(" "),i("tm-code-block",{staticClass:"codeblock",attrs:{language:"go",base64:"dHlwZSBEdXBsaWNhdGVWb3RlRXZpZGVuY2Ugc3RydWN0IHsKICBWb3RlQSAqVm90ZQogIFZvdGVCICpWb3RlCgogIC8vIGFiY2kgc3BlY2lmaWMgaW5mb3JtYXRpb24KICBUb3RhbFZvdGluZ1Bvd2VyIGludDY0CiAgVmFsaWRhdG9yUG93ZXIgICBpbnQ2NAogIFRpbWVzdGFtcCAgICAgICAgdGltZS5UaW1lCn0KCnR5cGUgTGlnaHRDbGllbnRBdHRhY2tFdmlkZW5jZSBzdHJ1Y3QgewogIENvbmZsaWN0aW5nQmxvY2sgKkxpZ2h0QmxvY2sKICBDb21tb25IZWlnaHQgICAgIGludDY0CgogIC8vIGFiY2kgc3BlY2lmaWMgaW5mb3JtYXRpb24KICBCeXphbnRpbmVWYWxpZGF0b3JzIFtdKlZhbGlkYXRvcgogIFRvdGFsVm90aW5nUG93ZXIgICAgaW50NjQgICAgICAgCiAgVGltZXN0YW1wICAgICAgICAgICB0aW1lLlRpbWUgCn0K"}}),e._v(" "),i("p",[e._v("These ABCI specific fields don't affect validity of the evidence itself but must\nbe consistent amongst nodes and agreed upon on chain. If evidence with the\nincorrect abci information is sent, a node will create new evidence from it and\nreplace the ABCI fields with the correct information.")])],1)}),[],!1,null,null,null);t.default=a.exports}}]);