import type { NextPage } from 'next'
import HeaderResponsive from '../../components/header'


const ModuleList: NextPage = () => {
  return (
    <>
      <HeaderResponsive links={[
        {
          label: "modules",
          link: "/"
        }
      ]}/>
      <h1>Hello Eve</h1>
    </>
  )
}

export default ModuleList