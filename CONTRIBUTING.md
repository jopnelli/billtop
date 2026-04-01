# Contributing to billtop

## Setup

```bash
git clone https://github.com/jopnelli/billtop.git
cd billtop
bun install
bun run test
bun run lint
```

## Project Structure

```
packages/
  core/   Shared logic: BigQuery queries, types, config
  cli/    CLI commands
  web/    Astro dashboard (not yet built)
```

## Making Changes

1. Fork and create a branch from `main`
2. Make your changes
3. Ensure `bun run lint && bun run typecheck && bun run test` pass
4. Open a pull request

### Commit Messages

[Conventional Commits](https://www.conventionalcommits.org/):

```
feat(report): add --format json flag
fix(bigquery): handle empty credits array
test(config): add validation edge cases
```

### Code Style

- TypeScript strict mode. No `any`.
- Biome for lint/format (`bun run lint:fix`).
- Tests alongside source (`foo.test.ts` next to `foo.ts`).
- Core logic in `packages/core/`. CLI and web are thin wrappers.
