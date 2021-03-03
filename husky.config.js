const tasks = (arr) => arr.join(' && ');

module.exports = {
  hooks: {
    'pre-commit': tasks([
      'echo notipie-ui lint:',
      'cd notipie-ui',
      'yarn lint',
      'cd ..',
      'echo notipie-ui-components lint:',
      'cd notipie-ui-components',
      'yarn lint',
      'cd ..',
    ]),
    'commit-msg': 'commitlint -E HUSKY_GIT_PARAMS',
  },
};
