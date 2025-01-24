# Torram Staking Prototype: Approach Document

## **Objective**

The goal is to create a prototype for staking Runes (or testnet BTC) on a Cosmos SDK chain named `Torram`. The state and data availability of staked assets should remain synchronized between the Cosmos SDK chain and the Bitcoin network.

## **Approach**

### **1. Setting Up the Cosmos SDK Blockchain**

- **Initialization:**
  - Used Starport to scaffold a new Cosmos SDK blockchain named `Torram`.
  - Set up a basic application structure with modules for custom logic.
- **Folder Structure:**
  - The `runestaking` module was created under `x/runestaking` to handle staking and unstaking functionality.

### **2. Defining Custom Message Types**

- **Stake Data Structure:**
  ```go
  type Stake struct {
      Address string `json:"address"`
      Amount  string `json:"amount"`
      Asset   string `json:"asset"` // BTC or Runes
  }
  ```
- **Message for Unstaking:**
  Defined `MsgUnstake` to represent unstaking requests:

  ```go
  type MsgUnstake struct {
      Creator string `json:"creator" yaml:"creator"`
      Amount  string `json:"amount" yaml:"amount"`
      Asset   string `json:"asset" yaml:"asset"`
  }

  func (msg MsgUnstake) GetSigners() []sdk.AccAddress {
      creatorAddr, err := sdk.AccAddressFromBech32(msg.Creator)
      if err != nil {
          panic("invalid address format")
      }
      return []sdk.AccAddress{creatorAddr}
  }
  ```

- The `GetSigners` method ensures that the Cosmos SDK can identify who is signing the transaction.

### **3. Handling Cosmos SDK and Bitcoin Network Synchronization**

- **State Synchronization:**
  - The staked assets’ state will be synchronized between the Torram chain and the Bitcoin network.
  - This ensures that when a staking or unstaking event occurs, both networks reflect the transaction accurately.

### **4. Debugging and Resolving Issues**

- **Challenges:**
  - Faced errors related to message registration and signer methods, such as:
    ```
    no cosmos.msg.v1.signer option found for message torram.runestaking.MsgUnstake
    ```
  - Fixed these by implementing the `DefineCustomGetSigners` method to properly register and identify custom message types.
- **Solution:**
  - Verified the protobuf definitions and ensured that the Cosmos SDK runtime correctly recognized custom messages.

### **5. Next Steps**

- **Finalize the Staking Module:**
  - Complete the implementation of staking and unstaking logic.
  - Add validation, state changes, and events for better tracking.
- **Integrate with Bitcoin Network:**
  - Implement interaction with the Bitcoin network to record transactions.
  - Use Bitcoin libraries (e.g., `btcsuite`) for testnet integration.
- **Testing:**
  - Test the end-to-end flow of staking and unstaking on the Torram chain.
  - Verify synchronization between the Torram chain and Bitcoin network.
- **Documentation:**
  - Add comprehensive documentation and a README to guide developers.

### **6. Current Status**

- **Progress:**
  - `Torram` chain scaffolded.
  - Basic types and messages for staking defined.
  - Working on resolving message signer registration issues.
- **Next Focus:**
  - Resolve errors during chain serving.
  - Implement state synchronization logic.

