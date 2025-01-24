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
