import { Avatar, Header, Container, Group, Title } from '@mantine/core';

import SearchBar from './searchBar'

export default function TweetvizHeader() {
  return (
    <Header height={60} mb={120}>
      <Container style={{display:'flex',justifyContent: 'space-between', alignItems: 'center', height: '100%'}}>
        <Group>
          <Avatar size="md" src="logo.png" alt="Tweetviz Logo" />
          <Title order={1}>Tweetviz</Title>
        </Group>
        <SearchBar />
        <></>
      </Container>
    </Header>
  );
}