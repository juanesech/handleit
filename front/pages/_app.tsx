import { MantineProvider } from '@mantine/core';
import type { AppProps } from 'next/app'
import React from 'react';


function MyApp({ Component, pageProps }: AppProps) {
  return (
    <MantineProvider
        withGlobalStyles
        withNormalizeCSS
        theme={{
          /** Put your mantine theme override here */
          colorScheme: 'light',
          fontFamily: 'Open Sans, sans serif',
          spacing: { xs: 15, sm: 20, md: 25, lg: 30, xl: 40 }
        }}
      >
      <Component {...pageProps} />
    </MantineProvider>
  );
}

export default MyApp