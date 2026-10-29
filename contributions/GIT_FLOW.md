# Git Flow

Branching model of HTTP Security Lab, based on the project rules for branches, commits and pull requests.

## 1. Long-lived branches


| Branch | Purpose                       | Created from | Merges into |
| ------ | ----------------------------- | ------------ | ----------- |
| main   | Stable, always buildable code | n/a          | n/a         |
| core   | Project core                  | main         | main        |
| back   | Web application backend       | main         | main        |
| front  | Frontend                      | main         | main        |
| sync   | Proxy synchronization         | main         | main        |

## 2. Working branches

Name format: `direction_purpose`


| Part      | Value           | Meaning                               |
| --------- | --------------- | ------------------------------------- |
| direction | `front`         | Frontend                              |
| direction | `back`          | Backend                               |
| direction | `ml`            | Machine learning                      |
| purpose   | service/feature | Work on a service or feature (`auth`) |
| purpose   | `tests`         | Tests                                 |
| purpose   | `docs`          | Documentation                         |
| purpose   | `fix`           | Bug fix                               |
| purpose   | `ref`           | Refactoring                           |


| Example         | Created from   | Merges into |
| --------------- | -------------- | ----------- |
| `back_auth`     | `back`         | `back`      |
| `back_tests`    | `back`         | `back`      |
| `front_account` | `front`        | `front`     |
| `front_docs`    | `front`        | `front`     |
| `back_fix`      | `main`(hotfix) | `main`      |

## 3. Commit messages

Format: `[direction:type]: message`


| Type    | Meaning          | Example                                               |
| ------- | ---------------- | ----------------------------------------------------- |
| `dev`   | Development      | `[back:dev]: add connection to PostgreSQL database`   |
| `feat`  | New feature      | `[back:feat]: add auth for users by jwt`              |
| `fix`   | Bug fix          | `[front:fix]: fix logo redirect to the main page`     |
| `tests` | Tests            | `[back:tests]: add test for AuthService:RegisterUser` |
| `ref`   | Refactoring      | `[back:ref]: split AuthService into modules`          |
| `ci/cd` | GitHub workflows | `[back:ci/cd]: add build workflow`                    |

## 4. Pull requests

Title format: `[direction:type]: message`


| Type    | Meaning     | Example                                 |
| ------- | ----------- | --------------------------------------- |
| `feat`  | Feature     | `[back:feat]: add password hashing`     |
| `fix`   | Bug fix     | `[front:fix]: fix account page crash`   |
| `tests` | Tests       | `[back:tests]: add auth service tests`  |
| `ref`   | Refactoring | `[back:ref]: simplify session handling` |


| Target                             | Merge method | Commit message        |
| ---------------------------------- | ------------ | --------------------- |
| Working branch into long-lived one | Squash merge | PR title              |
| Long-lived branch into`main`       | Merge commit | Default merge message |

## 5. Workflow


| Step | Action                         | Command / rule                                           |
| ---- | ------------------------------ | -------------------------------------------------------- |
| 1    | Update the base branch         | `git switch back`then`git pull origin back`              |
| 2    | Create a working branch        | `git switch -c back_auth`                                |
| 3    | Commit                         | `git commit -m "[back:feat]: add auth for users by jwt"` |
| 4    | Sync with the base branch      | `git fetch origin`then`git rebase origin/back`           |
| 5    | Run tests locally              | Required before opening a PR                             |
| 6    | Push                           | `git push -u origin back_auth`                           |
| 7    | Open a PR into the base branch | Title:`[back:feat]: add password hashing`                |
| 8    | Wait for CI/CD                 | Must pass before merging                                 |
| 9    | Wait for code review           | Required before merging                                  |
| 10   | Merge and clean up             | Squash merge, then`git push origin --delete back_auth`   |

## 6. Release to main


| Step | Action                                          | Command / rule                               |
| ---- | ----------------------------------------------- | -------------------------------------------- |
| 1    | Make sure the branch is stable                  | Project builds, tests pass                   |
| 2    | Open a PR from the long-lived branch into`main` | Same checks: local tests, CI/CD, code review |
| 3    | Merge with a merge commit                       | Keeps the component history                  |
| 4    | Tag the release                                 | `git tag -a v0.2.0 -m "Release 0.2.0"`       |
| 5    | Push the tag                                    | `git push origin --tags`                     |

## 7. Hotfix


| Step | Action                                  | Command / rule                                |
| ---- | --------------------------------------- | --------------------------------------------- |
| 1    | Branch from`main`                       | `git switch main`then`git switch -c back_fix` |
| 2    | Fix and commit                          | `[back:fix]: message`                         |
| 3    | Open a PR into`main`                    | Local tests, CI/CD, code review               |
| 4    | Update the affected long-lived branches | `git switch back`then`git merge main`         |

## 8. Rules


| Rule                                                      | Applies to                   |
| --------------------------------------------------------- | ---------------------------- |
| Never push directly to`main`                              | `main`                       |
| `main`contains only a stable app version                  | `main`                       |
| Do not merge long-lived branches into each other directly | `core`,`back`,`front`,`sync` |
| Update long-lived branches from`main`before new work      | `core`,`back`,`front`,`sync` |
| One feature or fix per branch and per PR                  | Working branches             |
| Force-push only on your own working branches              | Working branches             |
| Tags use the form`vMAJOR.MINOR.PATCH`                     | Releases                     |
