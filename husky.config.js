const tasks = (arr) => arr.join(' && ');

module.exports = {
  hooks: {
    'pre-commit': tasks([
      'echo "ui > lint (yarn lint)"',
      'cd ui',
      'yarn lint',
      'cd ..',
      'echo "ui-components > lint (yarn lint)"',
      'cd ui-components',
      'yarn lint',
      'cd ..',
    ]),
    'commit-msg': 'commitlint -E HUSKY_GIT_PARAMS',
  },
};
