1. **Variables and Constants**:
   - **CamelCase**: Use camelCase for variable names (e.g., `myVariable`, `maxLength`).
   - **MixedCase**: Constants are typically written in MixedCase (e.g., `MaxLength`, `TimeoutDuration`).

2. **Functions**:
   - **CamelCase**: Function names should also be in camelCase (e.g., `calculateSum`, `printMessage`).
   - **Exported Functions**: If a function is exported (i.e., visible outside its package), it should start with an uppercase letter (e.g., `AddNumbers`).

3. **Structs and Types**:
   - **CamelCase**: Struct and type names are written in CamelCase (e.g., `Person`, `ServerConfig`).

4. **Packages**:
   - **Lowercase**: Package names should be short and all lowercase without underscores or mixed caps (e.g., `http`, `math`, `strings`).
   - **Singular**: Use singular form for package names unless the package represents a collection of something (e.g., `image`, `bytes`).

5. **Files**:
   - **Lowercase with underscores**: File names are usually all lowercase with underscores separating words if needed (e.g., `main.go`, `server_config.go`).

6. **General Rules**:
   - **Avoid Abbreviations**: Use full words instead of abbreviations unless the abbreviation is widely understood (e.g., `HTML`, `URL`).
   - **No Hungarian Notation**: Go avoids prefixes like `str` for strings or `p` for pointers.

These conventions help ensure that Go code is consistent, readable, and easy to maintain.