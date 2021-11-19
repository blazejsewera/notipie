import type { ReactElement, ReactNode } from 'react'

type EmptyObject = Record<string, never>

/**
 * Default children type in React. Every renderable type is included.
 */
export type DefaultChildren = ReactNode

/**
 * No props, a.k.a. () => </> component
 */
export type EmptyProps = EmptyObject

/**
 * No React Context
 */
export type EmptyContext = EmptyObject

/**
 * Functional Component (strongly typed)
 */
export interface FC<P = EmptyProps, Ctx = EmptyContext> {
  (props: P, context?: Ctx): ReactElement | null
  displayName?: string
}
