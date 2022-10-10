import type { NextPage } from 'next'
import HeaderResponsive from '../../components/header'
import ModulePreview from '../../components/card'
import { ModuleSummary } from '../../interfaces/modules'
import { GetModules } from './modules'
import { useEffect, useState } from 'react'

const ModuleList: NextPage = () => {

  const [modules, setModules] = useState<ModuleSummary[]>([]);

  useEffect( () => {
    const getModules = async () => {
      let mods = await GetModules();
      console.log(mods)
      setModules(mods);
    }

    getModules();

  }, [modules])

  return (
    <>
      <HeaderResponsive links={[
        {
          label: "modules",
          link: "/"
        }
      ]}/>
      {modules.map( (module: ModuleSummary) => {
        return(
        <ModulePreview 
          title={module.Name}
          description={`Source: ${module.Providers}`}
        />
        )
      })}
    </>
  )
}

export default ModuleList