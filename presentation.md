---
# This Markdown file is designed for use with
# https://github.com/maaslalani/slides
author: "Akshay Shah"
date: 2024-03-14
# Presented to CNCF TAG Network
---

# Connect: simple, robust Protobuf RPC

Connect is a cross-language framework for building strongly-typed RPC APIs.
It's usually paired with Protocol Buffer schemas.

We're applying to join the CNCF as a sandbox project.

We've already presented to TAG Network and TAG App Delivery, but Nic wanted me
to present Connect's wire protocol in more detail.

- Website: https://connectrpc.com
- GitHub: https://github.com/connectrpc
- Application: https://github.com/cncf/sandbox/issues/63
- Sponsor: https://buf.build

---

## Multi-protocol support

Connect servers use content-type negotiation to simultaneously support three
RPC protocols:

* gRPC,
* gRPC-Web, and
* Connect's own protocol.

Clients can switch protocols with a configuration toggle. Application code is
protocol-agnostic, so switching protocols is easy.

When using the gRPC protocol, Connect applications work seamlessly with the
rest of the ecosystem.

---

## A new protocol?

Connect's protocol has the same semantics as gRPC, but it's simpler on the
wire. It supports streaming, JSON and binary encodings, leading and trailing
metadata, structured errors, and everything else you'd expect.

It works over HTTP/1.1, HTTP/2, and HTTP/3. It works in browsers, proxies,
and serverless environments. It works in Ruby on Rails and Next.js.

---

## Demo

Let's explore these protocols using a Connect server I'm running on port 8080.

The server uses the default configuration, so it supports all three RPC
protocols. With each protocol, it supports JSON and binary Protobuf.

The server's API is defined in a Protobuf schema, but it's not that
interesting - let's jump right into the deep end.

To get a feel for these protocols and the problems they solve, we're going to
start by digging into gRPC.

---

## Tradeoffs

* ✅ HTTP/1.1 support, no trailers: usable everywhere.
* ✅ HTTP/3 support: ready for the future.
* ✅ Standard HTTP semantics: better support from tools, proxies, and frameworks.
* ❌ Two protocols in one: more work for implementers.

---

## Questions? 🙋🏽

- Website: https://connectrpc.com
- GitHub: https://github.com/connectrpc
- Application: https://github.com/cncf/sandbox/issues/63
- Sponsor: https://buf.build
