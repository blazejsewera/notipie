const { jsWithBabelESM: tsjPreset } = require('ts-jest/presets')

/** @type {import('@jest/types').Config.InitialOptions} */
const config = {
  testRegex: '(/_test/.*|(\\.|/)(test|spec))\\.(jsx?|tsx?)$',
  transform: {
    ...tsjPreset.transform,
  },
  moduleNameMapper: {
    '\\.(jpg|jpeg|png|gif|eot|otf|webp|svg|ttf|woff|woff2|mp4|webm|wav|mp3|m4a|aac|oga)$':
      '<rootDir>/src/mock/file.mock.ts',
    '\\.(css|pcss)$': '<rootDir>/src/mock/style.mock.ts',
  },
  moduleFileExtensions: ['ts', 'tsx', 'js', 'jsx', 'json', 'node'],
  testEnvironment: 'jsdom',
}

module.exports = config
