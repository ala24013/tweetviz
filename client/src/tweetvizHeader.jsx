import { Avatar, Header, Container, Group, Title } from '@mantine/core';

import SearchBar from './searchBar'

export const HEADER_HEIGHT = 60;

export default function TweetvizHeader(props) {
  return (
    <Header height={60} style={{ display: 'flex', justifyContent: 'space-evenly', alignItems: "center" }}>
      <Container style={{ display:'flex', justifyContent: 'space-between',
        alignItems: 'center', height: '100%', width: '100%', maxWidth: '100%'}}>
        <Group style={{ width: '11vw', margin: '5vw', display: 'flex', justifyContent: 'space-around' }}>
          <Avatar size="md" src="logo.png" alt="Tweetviz Logo" />
          <Title order={1} style={{color: 'white'}}>Tweetviz</Title>
        </Group>
        <SearchBar sendMessage={props.sendMessage} loading={props.loading} />
        <Container style={{ margin: '5vw', width: '11vw' }}></Container>
      </Container>
    </Header>
  );
}