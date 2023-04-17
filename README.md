# ProjectStarter

ProjectStarter is a command-line utility that helps you bootstrap project structures for various types of projects, such as Terraform, Python, and more. It creates a project directory, generates a directory structure based on the chosen project type and pattern, and can include CI/CD configurations.

## Features

- Supports multiple project types (e.g., Terraform, Python, etc.)
- Configurable project structure patterns (e.g., basic, advanced)
- Optional CI/CD configuration inclusion (e.g., GitHub Actions, CircleCI, Jenkins, TravisCI)
- Creates a tests directory at the project root level (optional)
- Creates a README.md file
- Creates a .gitignore file
- Creates a LICENSE file
- Creates a .gitingore file

## Installation

- Install Go if you haven't already.
- Clone this repository:

```bash
git clone https://github.com/yourusername/projectstarter.git
```

- Change to the cloned directory:

```bash
cd projectstarter
```

- Build the binary:

```bash
go build -o projectstarter
```

- (Optional) Add the binary to your system's PATH for easier access.

## Usage

To create a new project:

```bash
./projectstarter create -t project-type -n project-name [-p pattern] [-c cicd] [-s]
-t, --project-type: The type of project to create (e.g., terraform, python).
-n, --name: The name of the project.
-p, --pattern: (Optional) The project structure pattern to use. Defaults to "basic" if not provided.
-c, --cicd: (Optional) Add a CI/CD directory (options: actions, circle, jenkins, travis).
-s, --tests: (Optional) Create a tests directory at the root level of the project.
```

## Example

Create a basic Terraform project with a GitHub Actions CI/CD configuration, and tests directory, using the basic project structure pattern:

```bash
./projectstarter create -t terraform -n my-terraform-project -p basic -c actions -s
```

Output structure:

```bash
my-terraform-project
├── README.md
├── environments
│   ├── dev
│   └── prod
├── main.tf
├── modules
├── outputs.tf
├── provider.tf
├── terraform.tfvars
├── tests
└── variables.tf
```

## Contributing

Contributions are welcome! Please feel free to submit issues or pull requests.

## License

This project is licensed under the MIT License. See the LICENSE file for details.