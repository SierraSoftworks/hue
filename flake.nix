{
  description = "Control your Hue lights directly from your terminal.";

  outputs = { self, nixpkgs }:
    {
      devShells = nixpkgs.lib.genAttrs ["aarch64-darwin" "aarch64-linux" "x86_64-darwin" "x86_64-linux"] (system:
        let
          pkgs = nixpkgs.legacyPackages.${system};
          lib = pkgs.lib;
          stdenv = pkgs.stdenv;
        in
        {
          default = stdenv.mkDerivation {
            name = "hue-shell-${system}";
            system = system;
            nativeBuildInputs = [
              pkgs.go
            ];
          };
        }
      );

      packages = nixpkgs.lib.genAttrs ["aarch64-darwin" "aarch64-linux" "x86_64-darwin" "x86_64-linux"] (system:
        let
          pkgs = nixpkgs.legacyPackages.${system};
          lib = pkgs.lib;
          stdenv = pkgs.stdenv;
        in
        {
          hue = pkgs.buildGoModule rec {
            name = "hue-${system}";
            pname = "hue";

            src = pkgs.nix-gitignore.gitignoreSourcePure ''
            /.github
            /target
            *.nix
            '' ./.;

            vendorHash = "sha256-WDqIOYEGr85eU6crclldnf4w9089+TsgbdnBAKhzRTM=";

            meta = with lib; {
              description = "Control your Hue lights directly from your terminal.";
              homepage = "https://github.com/SierraSoftworks/hue";
              license = licenses.mit;
              maintainers = [
                {
                  name = "Benjamin Pannell";
                  email = "contact@sierrasoftworks.com";
                }
              ];
            };
          };
          default = self.packages.${system}.hue;
        }
      );
    };
}
