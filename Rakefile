# helper functions...
# -----------------------------------------------------------------------------
def is_repo_clean?
  current_branch = `git rev-parse --abbrev-ref HEAD`.strip
  any_changes = `git status -s | wc -l`.strip.to_i
  if any_changes == 0
    true
  else
    false
  end
end

def command_exists?(cmd)
  cmd_check = `command -v #{cmd} > /dev/null 2>&1 && echo $?`.chomp
  cmd_check.length == 0 ? false : true
end

def current_version(lookup_file=".bumpversion.cfg")
  file = File.open(lookup_file, "r")
  data = file.read
  file.close
  match = /current_version = (\d+).(\d+).(\d+)/.match(data)
  "#{match[1]}.#{match[2]}.#{match[3]}"
end
# -----------------------------------------------------------------------------


# tasks
# -----------------------------------------------------------------------------
desc "Default task, show avaliable tasks"
task :default do
  system("rake -sT")
end

AVAILABLE_REVISIONS = ["major", "minor", "patch"]
task :bump, [:revision] do |_, args|
  args.with_defaults(revision: "patch")
  abort "bumpversion command not found..." unless command_exists?("bumpversion")
  abort "Please provide valid revision: #{AVAILABLE_REVISIONS.join(',')}" unless AVAILABLE_REVISIONS.include?(args.revision)

  system "bumpversion #{args.revision}"
end

desc "Run tests"
task :test, [:verbose] do |_, args|
  args.with_defaults(verbose: "")
  system "go test #{args.verbose} ./..."
end

ORGANIZATION_NAME = "promptapi"
REPO_NAME = "scraper-cli"

desc "Publish project with revision: #{AVAILABLE_REVISIONS.join(',')}, default: patch"
task :publish, [:revision] do |_, args|
  args.with_defaults(revision: "patch")
  abort "please commit your changes first..." unless is_repo_clean?
  Rake::Task["bump"].invoke(args.revision)
  
  current_git_tag = "v#{current_version}"
  puts "-> new version: #{current_git_tag}"
  puts "-> pushing tag #{current_git_tag} to remote..."
  system "git push origin #{current_git_tag}"
  
  current_branch = `git rev-parse --abbrev-ref HEAD`.strip
  puts "-> updating/pushing #{current_branch} branch"
  system "git push origin #{current_branch}"
  puts "-> all complete!"  
end

BUILD_OPTIONS = {
  windows: ['386', 'amd64'],
  linux: ['386', 'amd64'],
  macos: ['amd64'],
}

namespace :build do
  task :do, [:goos, :goarch, :target] do |_, args|
    rm_rf %w(build)
    system %{
      cd cmd/scraper-cli &&
      GOOS="#{args.goos}" GOARCH="#{args.goarch}" go build -o ../../build/#{args.target}
    }
    puts "-> build completed: build/#{args.target}"
  end
  
  desc "Build for Windows, available architecture(s): #{BUILD_OPTIONS[:windows].join(',')} default: #{BUILD_OPTIONS[:windows][1]}"
  task :windows, [:arch] do |_, args|
    args.with_defaults(arch: BUILD_OPTIONS[:windows][1])

    abort "Please provide valid arch: #{BUILD_OPTIONS[:windows].join(',')}" unless BUILD_OPTIONS[:windows].include?(args.arch)
    Rake::Task["build:do"].invoke("windows", args.arch, "scraper-cli-#{args.arch}.exe")
  end

  desc "Build for Linux, available architecture(s): #{BUILD_OPTIONS[:linux].join(',')} default: #{BUILD_OPTIONS[:linux][1]}"
  task :linux, [:arch] do |_, args|
    args.with_defaults(arch: BUILD_OPTIONS[:linux][1])

    abort "Please provide valid arch: #{BUILD_OPTIONS[:linux].join(',')}" unless BUILD_OPTIONS[:linux].include?(args.arch)
    Rake::Task["build:do"].invoke("linux", args.arch, "scraper-cli-linux-#{args.arch}")
  end

  desc "Build for macOS, available architecture(s): #{BUILD_OPTIONS[:macos].join(',')} default: #{BUILD_OPTIONS[:macos][0]}"
  task :macos, [:arch] do |_, args|
    args.with_defaults(arch: BUILD_OPTIONS[:macos][0])

    abort "Please provide valid arch: #{BUILD_OPTIONS[:macos].join(',')}" unless BUILD_OPTIONS[:macos].include?(args.arch)
    Rake::Task["build:do"].invoke("darwin", args.arch, "scraper-cli-macos-#{args.arch}")
  end

end
# -----------------------------------------------------------------------------