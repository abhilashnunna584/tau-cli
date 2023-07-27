package domainPrompts

import (
	structureSpec "github.com/taubyte/go-specs/structure"
	domainLib "github.com/taubyte/tau-cli/lib/domain"
	"github.com/taubyte/tau-cli/prompts"
	"github.com/urfave/cli/v2"
)

func New(ctx *cli.Context) (*structureSpec.Domain, error) {
	domain := &structureSpec.Domain{}

	taken, err := domainLib.List()
	if err != nil {
		return nil, err
	}

	domain.Name = prompts.GetOrRequireAUniqueName(ctx, NamePrompt, taken)
	domain.Description = prompts.GetOrAskForADescription(ctx)
	domain.Tags = prompts.GetOrAskForTags(ctx)

	if GetGeneratedFQDN(ctx) == true {
		generatedPrefix := GetGeneratedFQDNPrefix(ctx)

		domain.Fqdn, err = domainLib.NewGeneratedFQDN(generatedPrefix)
		if err != nil {
			return nil, err
		}
	} else {
		domain.Fqdn = GetOrRequireAnFQDN(ctx)
	}

	err = certificate(ctx, domain, true)
	if err != nil {
		return nil, err
	}

	return domain, nil
}
