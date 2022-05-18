(window.webpackJsonp=window.webpackJsonp||[]).push([[248],{822:function(e,t,n){"use strict";n.r(t);var o=n(1),s=Object(o.a)({},(function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("ContentSlotsDistributor",{attrs:{"slot-key":e.$parent.slotKey}},[n("h1",{attrs:{id:"block-sync"}},[n("a",{staticClass:"header-anchor",attrs:{href:"#block-sync"}},[e._v("#")]),e._v(" Block Sync")]),e._v(" "),n("p",[n("em",[e._v("Formerly known as Fast Sync")])]),e._v(" "),n("p",[e._v("In a proof of work blockchain, syncing with the chain is the same\nprocess as staying up-to-date with the consensus: download blocks, and\nlook for the one with the most total work. In proof-of-stake, the\nconsensus process is more complex, as it involves rounds of\ncommunication between the nodes to determine what block should be\ncommitted next. Using this process to sync up with the blockchain from\nscratch can take a very long time. It's much faster to just download\nblocks and check the merkle tree of validators than to run the real-time\nconsensus gossip protocol.")]),e._v(" "),n("h2",{attrs:{id:"using-block-sync"}},[n("a",{staticClass:"header-anchor",attrs:{href:"#using-block-sync"}},[e._v("#")]),e._v(" Using Block Sync")]),e._v(" "),n("p",[e._v("To support faster syncing, Tendermint offers a "),n("code",[e._v("blocksync")]),e._v(" mode, which\nis enabled by default, and can be toggled in the "),n("code",[e._v("config.toml")]),e._v(" or via\n"),n("code",[e._v("--blocksync.enable=false")]),e._v(".")]),e._v(" "),n("p",[e._v("In this mode, the Tendermint daemon will sync hundreds of times faster\nthan if it used the real-time consensus process. Once caught up, the\ndaemon will switch out of Block Sync and into the normal consensus mode.\nAfter running for some time, the node is considered "),n("code",[e._v("caught up")]),e._v(" if it\nhas at least one peer and it's height is at least as high as the max\nreported peer height. See "),n("a",{attrs:{href:"https://github.com/tendermint/tendermint/blob/b467515719e686e4678e6da4e102f32a491b85a0/blockchain/pool.go#L128",target:"_blank",rel:"noopener noreferrer"}},[e._v("the IsCaughtUp\nmethod"),n("OutboundLink")],1),e._v(".")]),e._v(" "),n("p",[e._v("Note: There are multiple versions of Block Sync. Please use v0 as the other versions are no longer supported.\nIf you would like to use a different version you can do so by changing the version in the "),n("code",[e._v("config.toml")]),e._v(":")]),e._v(" "),n("tm-code-block",{staticClass:"codeblock",attrs:{language:"toml",base64:"IyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIwojIyMgICAgICAgQmxvY2sgU3luYyBDb25maWd1cmF0aW9uIENvbm5lY3Rpb25zICAgICAgICMjIwojIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjCltibG9ja3N5bmNdCgojIElmIHRoaXMgbm9kZSBpcyBtYW55IGJsb2NrcyBiZWhpbmQgdGhlIHRpcCBvZiB0aGUgY2hhaW4sIEJsb2NrU3luYwojIGFsbG93cyB0aGVtIHRvIGNhdGNodXAgcXVpY2tseSBieSBkb3dubG9hZGluZyBibG9ja3MgaW4gcGFyYWxsZWwKIyBhbmQgdmVyaWZ5aW5nIHRoZWlyIGNvbW1pdHMKZW5hYmxlID0gdHJ1ZQoKIyBCbG9jayBTeW5jIHZlcnNpb24gdG8gdXNlOgojICAgMSkgJnF1b3Q7djAmcXVvdDsgKGRlZmF1bHQpIC0gdGhlIHN0YW5kYXJkIEJsb2NrIFN5bmMgaW1wbGVtZW50YXRpb24KIyAgIDIpICZxdW90O3YyJnF1b3Q7IC0gREVQUkVDQVRFRCwgcGxlYXNlIHVzZSB2MAp2ZXJzaW9uID0gJnF1b3Q7djAmcXVvdDsK"}}),e._v(" "),n("p",[e._v("If we're lagging sufficiently, we should go back to block syncing, but\nthis is an "),n("a",{attrs:{href:"https://github.com/tendermint/tendermint/issues/129",target:"_blank",rel:"noopener noreferrer"}},[e._v("open issue"),n("OutboundLink")],1),e._v(".")]),e._v(" "),n("h2",{attrs:{id:"the-block-sync-event"}},[n("a",{staticClass:"header-anchor",attrs:{href:"#the-block-sync-event"}},[e._v("#")]),e._v(" The Block Sync event")]),e._v(" "),n("p",[e._v("When the tendermint blockchain core launches, it might switch to the "),n("code",[e._v("block-sync")]),e._v("\nmode to catch up the states to the current network best height. the core will emits\na fast-sync event to expose the current status and the sync height. Once it catched\nthe network best height, it will switches to the state sync mechanism and then emit\nanother event for exposing the fast-sync "),n("code",[e._v("complete")]),e._v(" status and the state "),n("code",[e._v("height")]),e._v(".")]),e._v(" "),n("p",[e._v("The user can query the events by subscribing "),n("code",[e._v("EventQueryBlockSyncStatus")]),e._v("\nPlease check "),n("a",{attrs:{href:"https://pkg.go.dev/github.com/tendermint/tendermint/types?utm_source=godoc#pkg-constants",target:"_blank",rel:"noopener noreferrer"}},[e._v("types"),n("OutboundLink")],1),e._v(" for the details.")]),e._v(" "),n("h2",{attrs:{id:"implementation"}},[n("a",{staticClass:"header-anchor",attrs:{href:"#implementation"}},[e._v("#")]),e._v(" Implementation")]),e._v(" "),n("p",[e._v("To read more on the implamentation please see the "),n("RouterLink",{attrs:{to:"/tendermint-core/block-sync/reactor.html"}},[e._v("reactor doc")]),e._v(" and the "),n("RouterLink",{attrs:{to:"/tendermint-core/block-sync/implementation.html"}},[e._v("implementation doc")])],1)],1)}),[],!1,null,null,null);t.default=s.exports}}]);