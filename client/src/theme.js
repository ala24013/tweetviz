export const PRIMARY_COLOR = "#1D9BF0"

export const styles = {
  Title: (theme) => ({
    root: {
      color: PRIMARY_COLOR
    }
  })
}

export const theme = {
  colorScheme: 'dark',
  headings: {
    fontFamily: 'Roboto, sans-serif',
    sizes: {
      h1: { fontSize: 30 },
    }
  },
  colors: {
    primary: [ PRIMARY_COLOR ]
  },
  primaryColor: 'primary',
  loader: 'bars'
}