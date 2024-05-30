{
  inputs.nixpkgs.url = "github:NixOS/nixpkgs";

  outputs = {
    self,
    nixpkgs,
    ...
  }: let
    systems = ["x86_64-linux" "aarch64-linux" "x86_64-darwin" "aarch64-darwin"];
    forEachSystem = nixpkgs.lib.genAttrs systems;
    pkgsForEach = nixpkgs.legacyPackages;
  in {
    packages = forEachSystem (system: rec {
      zzz = pkgsForEach.${system}.callPackage ./default.nix {};

      default = zzz;
    });

    devShells = forEachSystem (system: {
      default = pkgsForEach.${system}.callPackage ./shell.nix {};
    });

    homeManagerModules.default = import ./hm-module.nix self;
  };
}
