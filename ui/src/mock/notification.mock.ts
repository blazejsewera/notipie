import type { Notification } from '../type/notification'
import githubIcon from './asset/github-icon.svg'

const body = `#12 add some new amazing functionality

Closes #10. Changes both in 'core' and 'ui'. Needs additional work with this and that.`

export const full: Notification = {
  appName: 'Github',
  title: 'New Pull Request',
  subtitle: '— notipie',
  body,
  timestamp: '2 hours ago',
  uuid: '1',
}

export const fullWithImage: Notification = {
  ...full,
  appImgUri: githubIcon,
}

export const partial: Notification = {
  appName: full.appName,
  title: full.title,
  body: full.body,
  timestamp: full.timestamp,
  uuid: '2',
}

export const minimal: Notification = {
  appName: full.appName,
  title: full.title,
  timestamp: full.timestamp,
  uuid: '3',
}
