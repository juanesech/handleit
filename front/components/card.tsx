import { createStyles, Paper, Text } from '@mantine/core';
import { FunctionComponent } from 'react';
import Link from 'next/link';

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

interface CardGradientProps {
  title: string;
  description: string;
}

const CardGradient: FunctionComponent<CardGradientProps> = ({ title, description }: CardGradientProps) => {
  const { classes } = useStyles();
  return (
    <Link href={`modules/${title}`}>
      <Paper withBorder radius="sm" className={classes.card}>
        <Text size="xl" weight={600} mt="s">
          {title}
        </Text>
        <Text size="sm" mt="s" color="dimmed">
          {description}
        </Text>
      </Paper>
    </Link>
  );
}

export default CardGradient