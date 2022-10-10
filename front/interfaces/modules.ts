export interface ModuleSummary {
  Name: string
  Providers: {
    source: string
  }[]
}

export interface Module {
  Id: string;
  Name: string;
  Variables: {
    name: string
    type: string
    description: string
    default: string
    required: boolean
  }[],
  Outputs: {
    name: string
    description: string
  }[],
  Providers: {
    source: string
    versionConstrains: string[]
  }[]
}