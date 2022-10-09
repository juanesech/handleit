export interface ModuleSummary {
  name: string
  providers: {
    source: string
  }[]
}

export interface Module {
  id: string;
  name: string;
  variables: {
    name: string
    type: string
    description: string
    default: string
    required: boolean
  },
  outputs: {
    name: string
    description: string
  },
  providers: {
    source: string
    versionConstrains: string[]
  }[]
}