export default (graphqlError, { store, error, redirect, route }) => {
  console.error({ graphqlError })

  const { networkError, message, gqlError, graphqlErrors } = graphqlError
  store.dispatch('application/alertError', {message: 'Failed, try again later.'})
  if (message === 'GraphQL error: Access denied') {
    console.log('attempting to dispatch message to global toast')
    store.dispatch('application/alertError', { message: 'Unable to load, required permissions not found' })
  }
}
