/**
 * Intl is an interface for an internationalization (i18l) tool
 * that can provide translations for page.
 */
export interface Intl {
  getTranslation: (text: string) => string
}

export const intl: Intl = { getTranslation: text => text }
