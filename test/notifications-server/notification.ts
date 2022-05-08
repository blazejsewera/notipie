import faker from '@faker-js/faker'

export type Notification = {
  appName: string
  timestamp: string
  appImgUri?: string
  title: string
  subtitle?: string
  body?: string
  id?: string
  read?: boolean
  extUri?: string
  readUri?: string
  archiveUri?: string
  relativeTime?: string
}

namespace util {
  export const generateArray = <T>(length: number, factory: () => T): T[] => Array.from([...Array(length)], factory)
  export const randBool = faker.datatype.boolean
  export const optional = <T>(element: T) => (randBool() ? element : undefined)
  export const alternative = <T>(first: T, second: T) => (randBool() ? first : second)
  export const titleCase = (s: string) => s.slice(0, 1).toUpperCase().concat(s.slice(1))
}

const exampleAppImgUri = 'https://fonts.gstatic.com/s/i/materialiconsround/star/v20/24px.svg'

const mockNotification = (appName: string, now?: boolean): Notification => ({
  appName,
  timestamp: now ? new Date().toISOString() : faker.date.recent().toISOString(),
  title: util.alternative(
    `${faker.word.verb()} ${faker.word.noun()}`,
    `${faker.word.verb()} ${faker.word.adjective()} ${faker.word.noun()}`,
  ),
  subtitle: util.optional(`${faker.word.adverb()} ${faker.word.verb()} ${faker.word.adjective()} ${faker.word.noun()}`),
  body: util.optional(faker.lorem.paragraph()),
  id: faker.random.alphaNumeric(45),
  read: util.randBool(),
  appImgUri: util.optional(exampleAppImgUri),
  archiveUri: faker.internet.url(),
  readUri: faker.internet.url(),
  extUri: faker.internet.url(),
})

const mockAppName = (): string =>
  util.alternative(
    `${util.titleCase(faker.word.adjective())} ${util.titleCase(faker.word.noun())}`,
    util.titleCase(faker.word.noun()),
  )

export const getMockNotification = (): Notification => mockNotification(mockAppName(), true)

export const getMockNotifications = (length: number): Notification[] => {
  const mockAppNamesLength = length < 3 ? 1 : Math.floor(length / 3)
  const mockAppNames = util.generateArray(mockAppNamesLength, mockAppName)
  return util.generateArray(length, () => mockNotification(faker.helpers.arrayElement(mockAppNames)))
}
