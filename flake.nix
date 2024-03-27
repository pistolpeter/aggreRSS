{
  description = "A simple Go package";

  # Nixpkgs / NixOS version to use.
  inputs = {
    nixpkgs.url = "nixpkgs/nixos-unstable";
    gomod2nix = {
      url = "github:tweag/gomod2nix";
      inputs.nixpkgs.follows = "nixpkgs";
    };
  };
  outputs = { self, nixpkgs, gomod2nix }:
    let

      # to work with older version of flakes
      lastModifiedDate = self.lastModifiedDate or self.lastModified or "19700101";

      # Generate a user-friendly version number.
      version = builtins.substring 0 8 lastModifiedDate;

      # System types to support.
      supportedSystems = [ "x86_64-linux" "x86_64-darwin" "aarch64-linux" "aarch64-darwin" ];

      # Helper function to generate an attrset '{ x86_64-linux = f "x86_64-linux"; ... }'.
      forAllSystems = nixpkgs.lib.genAttrs supportedSystems;

      # Nixpkgs instantiated for supported system types.
      nixpkgsFor = forAllSystems (system: import nixpkgs { 
        inherit system; 
        overlays = [ gomod2nix.overlays.default ];
      });

    in
    {
      packages = forAllSystems (system: 
      let
        pkgs = nixpkgsFor.${system};
      in 
      {
        default = pkgs.buildGoApplication {
          pname = "aggreRSS";
          version = "0.0.1";
          pwd = ./.;
          src = ./.;
          modules = ./gomod2nix.toml;
          go = pkgs.go;
        };
        
      });
      # Add dependencies that are only needed for development
      devShells = forAllSystems (system:
        let 
          pkgs = nixpkgsFor.${system};
        in
        {
          default = pkgs.mkShell {
            buildInputs = with pkgs; [ 
              go
              gopls 
              gotools 
              go-tools
              gomod2nix.packages.${system}.default
              postgresql
              pgadmin4
              sqlc
              goose
            ];
          };
        });
    };
}
