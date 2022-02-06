import { State } from '../../store/store'
import {
  fullWithHandlers,
  fullWithImageWithHandlers,
  minimalWithHandlers,
  otherAppWithHandlers,
  partialWithHandlers,
} from '../notification.mock'

const notificationsWithHandlers = [
  fullWithHandlers,
  fullWithImageWithHandlers,
  partialWithHandlers,
  minimalWithHandlers,
  otherAppWithHandlers,
]

export const mockState: State = {
  state: 'ok',
  notificationsWithHandlers: notificationsWithHandlers,
  isDarkMode: true,
}
