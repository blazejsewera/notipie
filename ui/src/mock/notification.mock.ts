import { Notification } from '../type/notification'
import { fiveSentenceLoremIpsum } from './asset/text/lipsum'
import githubIcon from './asset/icon/github-icon.svg'
import { NotificationCardHandlers } from '../type/handler'

const body = `#12 add some new amazing functionality

Closes #10. Changes both in 'core' and 'ui'. Needs additional work with this and that.`

export const mockNotificationCardHandlers: NotificationCardHandlers = {
  onArchive: () => {},
  onCheck: () => {},
  onExternal: () => {},
}

export const full: Notification = {
  appName: 'Github',
  title: 'New Pull Request',
  subtitle: '— notipie',
  body,
  timestamp: '2h ago',
  id: '0',
  handlers: mockNotificationCardHandlers,
}

export const fullWithImage: Notification = {
  ...full,
  id: '1',
  appImgUri: githubIcon,
  handlers: mockNotificationCardHandlers,
}

export const fullWithLoremIpsum: Notification = {
  ...full,
  id: '2',
  body: fiveSentenceLoremIpsum,
  handlers: mockNotificationCardHandlers,
}

export const partial: Notification = {
  appName: full.appName,
  title: full.title,
  body: full.body,
  timestamp: full.timestamp,
  id: '3',
  handlers: mockNotificationCardHandlers,
}

export const minimal: Notification = {
  appName: full.appName,
  title: full.title,
  timestamp: full.timestamp,
  id: '4',
  handlers: mockNotificationCardHandlers,
}

export const otherApp: Notification = {
  appName: 'Jenkins',
  title: 'Build succeeded',
  timestamp: full.timestamp,
  id: '5',
  handlers: mockNotificationCardHandlers,
}

export const read: Notification = {
  ...full,
  id: '6',
  read: true,
  handlers: mockNotificationCardHandlers,
}
