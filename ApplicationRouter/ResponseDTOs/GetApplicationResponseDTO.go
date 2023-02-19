package ResponseDTOs

import (
	"automation-suite/testUtils"
	"time"
)

type ApplicationResponseDTO struct {
	Errors []testUtils.Errors `json:"errors"`
	Result struct {
		Metadata struct {
			Name              string          `json:"name"`
			Namespace         string          `json:"namespace"`
			Uid               string          `json:"uid"`
			ResourceVersion   string          `json:"resourceVersion"`
			Generation        int             `json:"generation"`
			CreationTimestamp time.Time       `json:"creationTimestamp"`
			ManagedFields     []ManagedFields `json:"managedFields"`
		} `json:"metadata"`
		Spec struct {
			Source      Source `json:"source"`
			Destination struct {
				Server    string `json:"server"`
				Namespace string `json:"namespace"`
			} `json:"destination"`
			Project    string `json:"project"`
			SyncPolicy struct {
				Automated struct {
					Prune bool `json:"prune"`
				} `json:"automated"`
				Retry struct {
					Limit   int `json:"limit"`
					Backoff struct {
						Duration    string `json:"duration"`
						Factor      int    `json:"factor"`
						MaxDuration string `json:"maxDuration"`
					} `json:"backoff"`
				} `json:"retry"`
			} `json:"syncPolicy"`
		} `json:"spec"`
		Status struct {
			Resources []Resources `json:"resources"`
			Sync      struct {
				Status     string `json:"status"`
				ComparedTo struct {
					Source struct {
						RepoURL        string `json:"repoURL"`
						Path           string `json:"path"`
						TargetRevision string `json:"targetRevision"`
						Helm           struct {
							ValueFiles []string `json:"valueFiles"`
						} `json:"helm"`
					} `json:"source"`
					Destination struct {
						Server    string `json:"server"`
						Namespace string `json:"namespace"`
					} `json:"destination"`
				} `json:"comparedTo"`
				Revision string `json:"revision"`
			} `json:"sync"`
			Health struct {
				Status string `json:"status"`
			} `json:"health"`
			History []struct {
				Revision   string    `json:"revision"`
				DeployedAt time.Time `json:"deployedAt"`
				Id         int       `json:"id"`
				Source     struct {
					RepoURL        string `json:"repoURL"`
					Path           string `json:"path"`
					TargetRevision string `json:"targetRevision"`
					Helm           struct {
						ValueFiles []string `json:"valueFiles"`
					} `json:"helm"`
				} `json:"source"`
				DeployStartedAt time.Time `json:"deployStartedAt"`
			} `json:"history"`
			ReconciledAt   time.Time `json:"reconciledAt"`
			OperationState struct {
				Operation struct {
					Sync struct {
						Revision string `json:"revision"`
						Prune    bool   `json:"prune"`
					} `json:"sync"`
					InitiatedBy struct {
						Automated bool `json:"automated"`
					} `json:"initiatedBy"`
					Retry struct {
						Limit   int `json:"limit"`
						Backoff struct {
							Duration    string `json:"duration"`
							Factor      int    `json:"factor"`
							MaxDuration string `json:"maxDuration"`
						} `json:"backoff"`
					} `json:"retry"`
				} `json:"operation"`
				Phase      string `json:"phase"`
				Message    string `json:"message"`
				SyncResult struct {
					Resources []struct {
						Group     string `json:"group"`
						Version   string `json:"version"`
						Kind      string `json:"kind"`
						Namespace string `json:"namespace"`
						Name      string `json:"name"`
						Status    string `json:"status"`
						Message   string `json:"message"`
						HookPhase string `json:"hookPhase"`
						SyncPhase string `json:"syncPhase"`
					} `json:"resources"`
					Revision string `json:"revision"`
					Source   struct {
						RepoURL        string `json:"repoURL"`
						Path           string `json:"path"`
						TargetRevision string `json:"targetRevision"`
						Helm           struct {
							ValueFiles []string `json:"valueFiles"`
						} `json:"helm"`
					} `json:"source"`
				} `json:"syncResult"`
				StartedAt  time.Time `json:"startedAt"`
				FinishedAt time.Time `json:"finishedAt"`
			} `json:"operationState"`
			SourceType string `json:"sourceType"`
			Summary    struct {
				Images []string `json:"images"`
			} `json:"summary"`
		} `json:"status"`
	} `json:"result"`
}

type ManagedFields struct {
	Manager    string    `json:"manager"`
	Operation  string    `json:"operation"`
	ApiVersion string    `json:"apiVersion"`
	Time       time.Time `json:"time"`
	FieldsType string    `json:"fieldsType"`
	FieldsV1   struct {
		FStatus struct {
			Field1 struct {
			} `json:"."`
			FHealth struct {
				Field1 struct {
				} `json:"."`
				FStatus struct {
				} `json:"f:status"`
			} `json:"f:health"`
			FHistory struct {
			} `json:"f:history"`
			FOperationState struct {
				Field1 struct {
				} `json:"."`
				FFinishedAt struct {
				} `json:"f:finishedAt"`
				FMessage struct {
				} `json:"f:message"`
				FOperation struct {
					Field1 struct {
					} `json:"."`
					FInitiatedBy struct {
						Field1 struct {
						} `json:"."`
						FAutomated struct {
						} `json:"f:automated"`
					} `json:"f:initiatedBy"`
					FRetry struct {
						Field1 struct {
						} `json:"."`
						FBackoff struct {
							Field1 struct {
							} `json:"."`
							FDuration struct {
							} `json:"f:duration"`
							FFactor struct {
							} `json:"f:factor"`
							FMaxDuration struct {
							} `json:"f:maxDuration"`
						} `json:"f:backoff"`
						FLimit struct {
						} `json:"f:limit"`
					} `json:"f:retry"`
					FSync struct {
						Field1 struct {
						} `json:"."`
						FPrune struct {
						} `json:"f:prune"`
						FRevision struct {
						} `json:"f:revision"`
					} `json:"f:sync"`
				} `json:"f:operation"`
				FPhase struct {
				} `json:"f:phase"`
				FStartedAt struct {
				} `json:"f:startedAt"`
				FSyncResult struct {
					Field1 struct {
					} `json:"."`
					FResources struct {
					} `json:"f:resources"`
					FRevision struct {
					} `json:"f:revision"`
					FSource struct {
						Field1 struct {
						} `json:"."`
						FHelm struct {
							Field1 struct {
							} `json:"."`
							FValueFiles struct {
							} `json:"f:valueFiles"`
						} `json:"f:helm"`
						FPath struct {
						} `json:"f:path"`
						FRepoURL struct {
						} `json:"f:repoURL"`
						FTargetRevision struct {
						} `json:"f:targetRevision"`
					} `json:"f:source"`
				} `json:"f:syncResult"`
			} `json:"f:operationState"`
			FReconciledAt struct {
			} `json:"f:reconciledAt"`
			FResources struct {
			} `json:"f:resources"`
			FSourceType struct {
			} `json:"f:sourceType"`
			FSummary struct {
				Field1 struct {
				} `json:"."`
				FImages struct {
				} `json:"f:images"`
			} `json:"f:summary"`
			FSync struct {
				Field1 struct {
				} `json:"."`
				FComparedTo struct {
					Field1 struct {
					} `json:"."`
					FDestination struct {
						Field1 struct {
						} `json:"."`
						FNamespace struct {
						} `json:"f:namespace"`
						FServer struct {
						} `json:"f:server"`
					} `json:"f:destination"`
					FSource struct {
						Field1 struct {
						} `json:"."`
						FHelm struct {
							Field1 struct {
							} `json:"."`
							FValueFiles struct {
							} `json:"f:valueFiles"`
						} `json:"f:helm"`
						FPath struct {
						} `json:"f:path"`
						FRepoURL struct {
						} `json:"f:repoURL"`
						FTargetRevision struct {
						} `json:"f:targetRevision"`
					} `json:"f:source"`
				} `json:"f:comparedTo"`
				FRevision struct {
				} `json:"f:revision"`
				FStatus struct {
				} `json:"f:status"`
			} `json:"f:sync"`
		} `json:"f:status,omitempty"`
		FSpec struct {
			Field1 struct {
			} `json:"."`
			FDestination struct {
				Field1 struct {
				} `json:"."`
				FNamespace struct {
				} `json:"f:namespace"`
				FServer struct {
				} `json:"f:server"`
			} `json:"f:destination"`
			FProject struct {
			} `json:"f:project"`
			FSource struct {
				Field1 struct {
				} `json:"."`
				FHelm struct {
					Field1 struct {
					} `json:"."`
					FValueFiles struct {
					} `json:"f:valueFiles"`
				} `json:"f:helm"`
				FPath struct {
				} `json:"f:path"`
				FRepoURL struct {
				} `json:"f:repoURL"`
				FTargetRevision struct {
				} `json:"f:targetRevision"`
			} `json:"f:source"`
			FSyncPolicy struct {
				Field1 struct {
				} `json:"."`
				FAutomated struct {
					Field1 struct {
					} `json:"."`
					FPrune struct {
					} `json:"f:prune"`
				} `json:"f:automated"`
				FRetry struct {
					Field1 struct {
					} `json:"."`
					FBackoff struct {
						Field1 struct {
						} `json:"."`
						FDuration struct {
						} `json:"f:duration"`
						FFactor struct {
						} `json:"f:factor"`
						FMaxDuration struct {
						} `json:"f:maxDuration"`
					} `json:"f:backoff"`
					FLimit struct {
					} `json:"f:limit"`
				} `json:"f:retry"`
			} `json:"f:syncPolicy"`
		} `json:"f:spec,omitempty"`
	} `json:"fieldsV1"`
}

type Source struct {
	RepoURL        string `json:"repoURL"`
	Path           string `json:"path"`
	TargetRevision string `json:"targetRevision"`
	Helm           struct {
		ValueFiles []string `json:"valueFiles"`
	} `json:"helm"`
}

type Resources struct {
	Version   string `json:"version"`
	Kind      string `json:"kind"`
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
	Status    string `json:"status"`
	Health    struct {
		Status  string `json:"status"`
		Message string `json:"message,omitempty"`
	} `json:"health"`
	Group string `json:"group,omitempty"`
}
