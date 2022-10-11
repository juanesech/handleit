import type { NextPage } from 'next'
import HeaderResponsive from '../../components/header'
import ModulePreview from '../../components/card'
import { ModuleSummary } from '../../interfaces/modules'
import { GetModules } from './modules'
import { useEffect, useState } from 'react'
import { Container } from '@mantine/core'

const ModuleList: NextPage = () => {

  const [modules, setModules] = useState<ModuleSummary[]>([]);

  useEffect( () => {
    const getModules = async () => {
      let mods = await GetModules();
      console.log(mods)
      setModules(mods);
    }

    getModules();

  }, [])

  return (
    <>
      <HeaderResponsive links={[
        {
          label: "modules",
          link: "modules/"
        }
      ]}/>
      <Container my="xl">
        {modules.map( (module: ModuleSummary) => {
          return(
          <ModulePreview 
            title={module.Name}
            description={`Source: ${module.Providers}`}
          />
          )
        })}
      </Container>
    </>
  )
}

export default ModuleList