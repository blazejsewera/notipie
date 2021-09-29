export const cx = (...classNames: string[]): string =>
  classNames.length === 0 ? '' : classNames.reduce((previous, current) => `${previous} ${current}`)
