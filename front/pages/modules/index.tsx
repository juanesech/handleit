import type { NextPage } from 'next'
import HeaderResponsive from '../../components/header'
import ModulePreview from '../../components/card'
import {ModuleSummary} from '../../interfaces/modules'
import { ListModules } from './modules'


let moduleList: Array<ModuleSummary> = ListModules()

const ModuleList: NextPage = () => {
  return (
    <>
      <HeaderResponsive links={[
        {
          label: "modules",
          link: "/"
        }
      ]}/>
      {moduleList.map( module => {
        return(
        <ModulePreview 
          title={module.name}
          description={`Source: ${module.providers[0].source}`}
        />
        )
      })}
    </>
  )
}

export default ModuleList