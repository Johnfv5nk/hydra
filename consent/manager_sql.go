package consent

import (
	"context"
	"github.com/ory/hydra/driver/config"
)

// Assuming ManagerSQL struct exists
func (m *ManagerSQL) RevokeUserLoginSession(ctx context.Context, subject string) error {
	// Existing DB revocation logic...
	err := m.db.RevokeUserLoginSession(ctx, subject)
	if err != nil {
		return err
	}
	// Invalidate cache
	return m.cache.Delete(ctx, "session:" + subject)
}

func (m *ManagerSQL) RevokeUserConsentSession(ctx context.Context, subject string) error {
	// Existing DB revocation logic...
	err := m.db.RevokeUserConsentSession(ctx, subject)
	if err != nil {
		return err
	}
	// Invalidate cache
	return m.cache.Delete(ctx, "consent:" + subject)
}