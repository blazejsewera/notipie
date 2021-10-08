import type { ReactElement, ReactNode } from 'react'

type EmptyObject = Record<string, never>
type Children = ReactNode

/**
 * No props, a.k.a. () => </>, or ({ children }) => </> component
 */
export type EmptyProps = EmptyObject

/**
 * No React Context
 */
export type EmptyContext = EmptyObject

/**
 * Functional Component without children
 */
export interface FC<P = EmptyProps, Ctx = EmptyContext> {
  (props: P, context?: Ctx): ReactElement | null
  displayName?: string
}

/**
 * Functional Component with children
 */
export interface FCC<C = Children, Ctx = EmptyContext> {
  (props: { children?: C }, context?: Ctx): ReactElement | null
  displayName?: string
}

/**
 * Functional Component with props and children
 */
export interface FCPC<P = EmptyProps, C = Children, Ctx = EmptyContext> {
  (props: P & { children?: C }, context?: Ctx): ReactElement | null
  displayName?: string
}
