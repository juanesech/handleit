import type { NextPage } from 'next'
import HeaderResponsive from '../../../components/header'
import CardGradient from '../../../components/card'
import { useRouter } from 'next/router'


const Module: NextPage = () => {
  const router = useRouter()
  return (
    <>
      <HeaderResponsive links={[
        {
          label: "modules",
          link: "/"
        },
        {
          label: "test",
          link: "/home"
        }
      ]}/>
      <CardGradient 
        title={String(router.query.module)}
        description={"Hello Eve"}
      />
    </>
  )
}

export default Module