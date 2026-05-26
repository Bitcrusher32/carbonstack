# OpenMLS Sidecar Test Suite Split Plan v0

Status: Go test-suite maintainability plan
Component: CarbonStackComms / internal/protocol OpenMLS sidecar tests
Phase: v0.2.42 planning for v0.2.45 implementation
Previous docs:
- docs/86-openmls-sidecar-phase2d-mainline-closure-result-v0.md
- docs/88-openmls-sidecar-maintainability-promotion-plan-v0.md

## 1. Purpose

This document plans the split of the large Go OpenMLS sidecar contract test file.

Current file:

    carbonstack-comms/internal/protocol/openmls_sidecar_provider_info_test.go

Current approximate size:

    about 2,938 lines

The test coverage is valuable. The file shape is now the maintainability problem.

## 2. Current file responsibilities

The current file covers:

- provider-info;
- unsupported command behavior;
- identity-create;
- identity-status;
- public-bundle export;
- public-bundle artifact export;
- conversation-create;
- conversation-load-check;
- conversation-add-member;
- conversation-join;
- message-protect;
- message-open;
- two-message continuity;
- out-of-order two-message open;
- duplicate/replay rejection;
- corrupt/truncated artifact rejection;
- wrong-device rejection;
- wrong-conversation rejection;
- bidirectional message flow;
- shared sidecar runner;
- shared envelope structs;
- shared assertions;
- shared setup helpers.

All of this should stay. It should be split by ownership.

## 3. Proposed split

Recommended files:

    openmls_sidecar_helpers_test.go
    openmls_sidecar_provider_info_test.go
    openmls_sidecar_identity_test.go
    openmls_sidecar_public_bundle_test.go
    openmls_sidecar_conversation_test.go
    openmls_sidecar_message_test.go
    openmls_sidecar_message_negative_test.go

## 4. File ownership

openmls_sidecar_helpers_test.go:

- runOpenMLSSidecar;
- sidecar path resolution;
- removeOpenMLSSidecarState;
- envelope structs;
- parseSidecarEnvelope;
- assertSidecarError;
- assertNoSecretMaterialInStdout;
- assertFileExists;
- assertExitCode;
- shared lifecycle setup helpers:
  - setupOpenMLSTwoMemberConversation;
  - protectOpenMLSSidecarMessage;
  - openOpenMLSSidecarMessage;
- message success assertion helpers.

openmls_sidecar_provider_info_test.go:

- provider-info output shape;
- capability list;
- unsupported command list;
- unsupported command envelope tests;
- command surface claims.

openmls_sidecar_identity_test.go:

- identity-create success;
- identity-create duplicate;
- identity invalid label;
- identity-status success;
- identity-status missing state;
- identity private material redaction checks.

openmls_sidecar_public_bundle_test.go:

- public-bundle-export summary mode;
- public-bundle-export write-artifact mode;
- public bundle manifest checks;
- KeyPackage artifact path/size/hash checks;
- provider storage persistence required for Welcome consumption.

openmls_sidecar_conversation_test.go:

- conversation-create;
- conversation-load-check;
- conversation-add-member Welcome export;
- conversation-join Welcome consume;
- device-scoped conversation path assertions.

openmls_sidecar_message_test.go:

- message-protect/open one-way;
- message-protect/open two sequential messages;
- out-of-order two-message open;
- bidirectional Alice/Bob message flow.

openmls_sidecar_message_negative_test.go:

- duplicate/replay open rejection;
- corrupt/truncated artifact rejection;
- wrong-device open rejection;
- wrong-conversation open rejection;
- future missing-artifact or invalid-label tests.

## 5. Test name preservation

Prefer keeping existing test names unless a rename is necessary.

Current important tests include:

- TestOpenMLSSidecarProviderInfoCommand;
- TestOpenMLSSidecarUnsupportedCommandEnvelope;
- TestOpenMLSSidecarConversationCreate;
- TestOpenMLSSidecarConversationLoadCheck;
- TestOpenMLSSidecarConversationAddMemberWelcomeExport;
- TestOpenMLSSidecarConversationJoinWelcomeConsume;
- TestOpenMLSSidecarMessageProtectOpenOneWay;
- TestOpenMLSSidecarMessageProtectOpenTwoSequentialMessages;
- TestOpenMLSSidecarMessageOpenOutOfOrderTwoMessageDelivery;
- TestOpenMLSSidecarMessageOpenDuplicateRejected;
- TestOpenMLSSidecarMessageOpenCorruptArtifactRejected;
- TestOpenMLSSidecarMessageOpenWrongDeviceRejected;
- TestOpenMLSSidecarMessageOpenWrongConversationRejected;
- TestOpenMLSSidecarMessageProtectOpenBidirectional.

## 6. Test target transition

After v0.2.43 promoted sidecar scaffold exists, tests should move from targeting the research sidecar to targeting the promoted sidecar.

The research sidecar should not be deleted.

Recommended policy:

- perform a one-time equivalence proof;
- then make the promoted sidecar the active test target;
- keep research sidecar as known-good reference, not permanent duplicate test target.

## 7. Validation target

After test split:

- gofmt all split test files;
- targeted Phase 2D closure subset passes;
- go test -p 1 ./internal/protocol passes;
- go test -p 1 ./... passes;
- Rust artifact guard passes.

## 8. Non-goals

Do not weaken test coverage.

Do not delete negative tests.

Do not remove secret-redaction assertions.

Do not collapse wrong-device and wrong-conversation tests into vague generic tests.

Do not rewrite behavior while splitting tests.

Do not start Cypher routing in the test split rung.
