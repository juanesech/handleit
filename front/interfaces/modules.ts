export interface ModuleSummary {
  Name: string
  Providers: {
    source: string
  }[]
}

export interface Module {
  ID: string;
  Name: string;
  Variables: {
    Name: string
    Type: string
    Description: string
    Default: string
    Required: boolean
  }[],
  Outputs: {
    Name: string
    Description: string
  }[],
  Providers: {
    source: string
    versionConstrains: string[]
  }[]
}