const tasks = (arr) => arr.join(' && ');

module.exports = {
  hooks: {
    'pre-commit': tasks(['cd notipie-ui', 'yarn lint', 'cd ..']),
    'commit-msg': 'commitlint -E HUSKY_GIT_PARAMS',
  },
};
