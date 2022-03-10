import c from '../../notipie.config.json'

type Config = {
  mode: 'dev' | 'prod'
}

export const config = c as Config
