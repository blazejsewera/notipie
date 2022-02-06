import { createStore } from 'redux'
import { mockState } from './state.mock'

const reducer = (previousState = mockState) => previousState

export const mockStore = createStore(reducer)
