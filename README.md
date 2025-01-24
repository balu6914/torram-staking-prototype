# Torram Staking Prototype

This repository contains a prototype for staking Runes (or BTC as a placeholder) on a custom Cosmos SDK blockchain called **Torram**. The goal is to create a decentralized application that allows staking and unstaking of assets (Runes/BTC) on the Cosmos SDK chain, with transactional information being visible on both the Cosmos SDK chain and the Bitcoin network.

---

## Table of Contents

- [Overview](#overview)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installing](#installing)
- [Running the Blockchain](#running-the-blockchain)
- [Project Structure](#project-structure)
- [Usage](#usage)
- [Dependencies](#dependencies)
- [License](#license)

---

## Overview

The **Torram Staking Prototype** aims to:

- Facilitate **staking and unstaking** of assets (Runes/BTC).
- Ensure **real-time synchronization** of asset data between the Bitcoin and Cosmos SDK networks.
- Use custom Cosmos SDK message types for transaction handling (e.g., `MsgUnstake`).

---

## Getting Started

### Prerequisites

Before starting, ensure you have the following installed:

- **Go** (v1.23.5 or higher)
- **Cosmos SDK** (v0.50.11 via Starport)
- **Starport CLI** (v28.7.0)
- **Node.js** (for potential API/front-end integration)

### Installing

1. Clone this repository:
   ```bash
   git clone https://github.com/balu6914/torram-staking-prototype.git
   cd torram-staking-prototype
   ```

Install Cosmos SDK dependencies:

bash
Copy
Edit
cd cosmos-sdk
make install
Scaffold the Torram chain using Starport:

bash
Copy
Edit
cd ..
starport scaffold chain github.com/balu6914/torram
Modify the custom modules (runestaking) in the x/ directory to handle staking and unstaking.

Running the Blockchain
Start the chain locally using Starport:

bash
Copy
Edit
cd torram
starport chain serve
This will build and start the chain, enabling you to interact with it via the CLI.

Project Structure
perl
Copy
Edit
torram-staking-prototype/
├── cosmos-sdk/ # Cosmos SDK configurations
├── torram/ # Torram blockchain source code
│ ├── cmd/ # CLI commands
│ ├── x/ # Modules (e.g., runestaking)
│ ├── app/ # Application setup
│ └── ...
└── README.md # Project documentation
Usage
Interact with the chain using Cosmos SDK CLI or custom transaction messages.

Example CLI Command:
bash
Copy
Edit
cosmosd tx runestaking unstake [AMOUNT] [ASSET] --from [SENDER]
[AMOUNT]: Amount to unstake (e.g., 100).
[ASSET]: Type of asset (e.g., BTC or Runes).
[SENDER]: Address performing the transaction.
Dependencies
Cosmos SDK (v0.50.11)
Go (v1.23.5)
Starport CLI (v28.7.0)
Bitcoin Go Library (for BTC integration)
Refer to:

Cosmos SDK Documentation
Starport Documentation

License
This project is licensed under the MIT License. See the LICENSE file for more details.
