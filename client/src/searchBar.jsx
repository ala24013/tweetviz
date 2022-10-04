import React from "react"
import { TextInput, Button, Box, ActionIcon } from "@mantine/core"
import { IconSearch } from "@tabler/icons"

DEFAULT_SEARCH = "Ian"

export default function SearchBar(props) {
    const [searchValue, setSearchValue] = useState(DEFAULT_SEARCH)

    return (
        <Box>
            <TextInput
                value={searchValue}
                onChange={(event) => setSearchValue(event.currentTarget.value)}
                radius="xl"
                size="lg"
                error={/^[0-9a-zA-Z]{0,50}$/.test(searchValue) ? null : "ASCII characters and Numbers only!"}
                rightSection={
                    <Button variant="subtle">
                        <ActionIcon variant="subtle">
                            <IconSearch />
                        </ActionIcon>
                    </Button>
                }
            />
        </Box>
    )
}