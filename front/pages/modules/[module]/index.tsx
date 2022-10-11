import type { NextPage } from 'next';
import HeaderResponsive from '../../../components/header';
import { Container, Paper, Text, createStyles } from '@mantine/core';
import { Module } from '../../../interfaces/modules';
import { useEffect, useState } from 'react';
import { GetModule } from '../modules';
import { useRouter } from 'next/router';

const useStyles = createStyles((theme) => ({
  card: {
    position: 'relative',
    overflow: 'hidden',
   
    padding: theme.spacing.xl,
    paddingLeft: theme.spacing.xl * 2,

    '&::before': {
      content: '""',
      position: 'absolute',
      top: 0,
      bottom: 0,
      left: 0,
      width: 6,
      backgroundImage: theme.fn.linearGradient(0, theme.colors.pink[4], theme.colors.orange[6]),
    },
  },
}));

const Module: NextPage = () => {

  const router = useRouter();
  const [module, setModule] = useState<Module>({
    Name: "",
    ID: "",
    Variables: [{
      Name: "",
      Description: "",
      Default: "",
      Required: false,
      Type: ""
    }],
    Outputs: [{
      Name: "",
      Description: ""
    }],
    Providers: [{
      source: "",
      versionConstrains: [""]
    }], 
  });

  useEffect( () => {
    const getModules = async () => {
      let mod = await GetModule(router.query.module)
      console.log(mod)
      setModule(mod);
    }

    getModules();

  }, [])


  const { classes } = useStyles();
  return (
    <>
      <HeaderResponsive links={[
        {
          label: "modules",
          link: "/modules"
        }
      ]}/>
    <Container my="xl">
      <Paper withBorder radius="sm" className={classes.card}>
        <Text size="lg" weight={600} mt="s">
          {module.Name}
        </Text>
        <Text size="sm" mt="s" color="dimmed">
          {module.ID}
        </Text>
        <Text size="sm" mt="s" color="dimmed">
          {"Variables"}
        </Text>
          {module.Variables.map((variable) => {
              return (
                <Container my={"xs"}>
                  <Text size="sm" mt="s" color="dimmed">
                    {variable.Name}
                  </Text>
                  <Container my={"xxs"}>
                    <Text size="xs" mt="s" color="dimmed">
                      {`Description: ${variable.Description}`}
                    </Text>
                    <Text size="xs" mt="s" color="dimmed">
                      {`Required: ${variable.Required}`}         
                    </Text>
                    <Text size="xs" mt="s" color="dimmed">
                      {`Default: ${variable.Default}`}         
                    </Text>
                    <Text size="xs" mt="s" color="dimmed">
                      {`Type: ${variable.Type}`}         
                    </Text>
                  </Container>
                </Container>
              )
          })}
      </Paper>
    </Container>
    </>
  )
}

export default Module