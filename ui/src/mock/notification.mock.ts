import type { Notification } from '../type/notification'
import { fiveSentenceLoremIpsum } from './asset/text/lipsum'
import githubIcon from './asset/icon/github-icon.svg'
import { NotificationCardHandlers } from '../type/handler'

const body = `#12 add some new amazing functionality

Closes #10. Changes both in 'core' and 'ui'. Needs additional work with this and that.`

const handlers: NotificationCardHandlers = {
  onArchive: () => {},
  onCheck: () => {},
  onSettings: () => {},
}

export const full: Notification = {
  appName: 'Github',
  title: 'New Pull Request',
  subtitle: '— notipie',
  body,
  timestamp: '2 hours ago',
  id: '0',
  handlers,
}

export const fullWithImage: Notification = {
  ...full,
  id: '1',
  appImgUri: githubIcon,
}

export const fullWithLoremIpsum: Notification = {
  ...full,
  id: '2',
  body: fiveSentenceLoremIpsum,
}

export const partial: Notification = {
  appName: full.appName,
  title: full.title,
  body: full.body,
  timestamp: full.timestamp,
  id: '3',
  handlers,
}

export const minimal: Notification = {
  appName: full.appName,
  title: full.title,
  timestamp: full.timestamp,
  id: '4',
  handlers,
}
