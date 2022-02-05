type RecursiveArray<T> = Array<RecursiveArray<T> | T>

type RecursiveFlatten = (arr: RecursiveArray<string>) => Array<string>
const flatten: RecursiveFlatten = arr => arr.flat(10) as Array<string>

export const cx = (...classNames: RecursiveArray<string>): string =>
  classNames.length === 0 ? '' : flatten(classNames).reduce((previous, current) => `${previous} ${current}`)
