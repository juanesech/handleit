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
    cursor: 'pointer',
    overflow: 'hidden',
    transition: 'transform 150ms ease, box-shadow 100ms ease',
    padding: theme.spacing.xl,
    paddingLeft: theme.spacing.xl * 2,

    '&:hover': {
      boxShadow: theme.shadows.md,
      transform: 'scale(1.02)',
    },

    '&::before': {
      content: '""',
      position: 'absolute',
      top: 0,
      bottom: 0,
      left: 0,
      width: 6,
      backgroundImage: theme.fn.linearGradient(0, theme.colors.pink[6], theme.colors.orange[6]),
    },
  },
}));

const Module: NextPage = () => {

  const router = useRouter();
  const [module, setModule] = useState<Module>({
    Name: "",
    Id: "",
    Variables: [{
      name: "",
      description: "",
      default: "",
      required: false,
      type: ""
    }],
    Outputs: [{
      name: "",
      description: ""
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

  }, [module])


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
          {module.Id}
        </Text>
      </Paper>
    </Container>
    </>
  )
}

export default Module