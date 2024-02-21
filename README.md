# License Wizard

License Wizard is a command-line tool written in Go that generates license files for your git projects. It provides a simple interface for users to specify the license type, author, project name, and output directory, and then generates the corresponding license file accordingly.

## Features

- Supports generation of all available license files for use on GitHub.
- Allows users to specify author name, project name, and output directory for the generated license file.
- The data is dynamically generated

## Usage

To generate a license file, run the following command:

```bash
licensewizard <license-type>
```

Replace `<license-type>` with one of the supported license types, such as `MIT`, `BSD_2.0`, `Apache_2.0` (If the license type has a version the number is prefaced with a '_' )

The program will then prompt you to enter the author name, project name, and output directory for the generated license file.

Example:

```bash
$ license-generator MIT
Author: John Doe
Project name: MyProject
Output Directory: /path/to/output
```

The generated license file will be saved in the specified output directory with the filename `LICENSE`.

## Supported License Types

License Generator supports the following license types:

- MIT
- BSD 2.0
- BSD 3.0
- Apache 2.0
- AGPL 3.0
- GPL 3.0
- GPL 2.0
- MPL 2.0
- CCZERO 1.0
- ECLIPSE 2.0
- LGPL 2.1
- UNLICENSE
- BOOST


## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
