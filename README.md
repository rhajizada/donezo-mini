# donezo-mini

**donezo-mini** is a streamlined version of the
[**donezo**](https://github.com/rhajizada/donezo.git) project, focused solely on
managing tasks through a text-based user interface (TUI) without web components.
The TUI interacts directly with a local SQLite database to manage boards and items.

## Features

- Task Management: Organize tasks into boards, each with individual items.
- TUI Interface: An interactive command-line UI built with Bubble Tea.
- SQLite Database: Data is stored locally in an SQLite database.
- Boards and Items: Create, update, delete, and list boards and items, with
  support for toggling item completion status.

## Installation

1. Clone the repository:

```bash
git clone https://github.com/rhajizada/donezo-mini.git
cd donezo-mini
```

2. Install the application:

```bash
make install
```

## Roadmap

- 0.1.0

  - Implement comprehensive test coverage for all modules.
  - Set up CI using GitHub Actions.
  - Automate releases and publish them via GitHub Actions.

- 0.2.0

  - Add support for custom app styling.
  - Introduce configuration options for app customization.

- 0.3.0
  - `Tags`
    - Enable items on boards to be tagged with any word.
    - Add aggregated tag views, accessible in both the service layer and the TUI.
