import { config } from '../index'

it('properly reads the config file', () => {
  expect(config).toHaveProperty('endpointConfig')
})
