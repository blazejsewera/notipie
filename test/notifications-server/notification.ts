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

const generateArray = <T>(length: number, factory: () => T): T[] => Array.from([...Array(length)], factory)

const optional = <T>(element: T) => (faker.datatype.boolean() ? element : undefined)

const exampleAppImgUri = 'https://fonts.gstatic.com/s/i/materialiconsround/star/v20/24px.svg'

const mockNotification = (appName: string): Notification => ({
  appName,
  timestamp: faker.date.recent().toISOString(),
  title: `${faker.word.verb()} ${faker.word.noun()}`,
  subtitle: optional(`${faker.word.adverb()} ${faker.word.verb()} ${faker.word.adjective()} ${faker.word.noun()}`),
  body: optional(faker.lorem.paragraph()),
  id: faker.random.alphaNumeric(45),
  read: faker.datatype.boolean(),
  appImgUri: optional(exampleAppImgUri),
  archiveUri: faker.internet.url(),
  readUri: faker.internet.url(),
  extUri: faker.internet.url(),
})

const mockAppName = (): string => faker.company.companyName()

export const getMockNotifications = (length: number): Notification[] => {
  const mockAppNamesLength = length < 3 ? 1 : Math.floor(length / 3)
  const mockAppNames = generateArray(mockAppNamesLength, mockAppName)
  return generateArray(length, () => mockNotification(faker.helpers.arrayElement(mockAppNames)))
}
