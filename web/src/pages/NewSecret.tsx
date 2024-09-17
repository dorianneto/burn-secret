import React from "react";
import {
  Button,
  Tooltip,
  useToast,
} from "@chakra-ui/react";
import { Card, CardBody } from "@chakra-ui/react";
import { CopyIcon } from "@chakra-ui/icons";

export const NewSecret = () => {
  const toast = useToast();

  return (
    <Card variant="elevated">
      <CardBody>
        <Tooltip
          label="Reminder: this link will be burned within 12 hours"
          placement="top"
        >
          <p id="secret-link" className="text-xl my-8">
            {`http://localhost:8080/secret/123/reveal`}
          </p>
        </Tooltip>
        <Button
            leftIcon={<CopyIcon />}
          colorScheme="orange"
          onClick={() => {
            const secretLink = document.getElementById("secret-link");
            navigator.clipboard.writeText(secretLink?.textContent || "");
            toast({
              title: "Secret link copied!",
              status: "success",
              duration: 6000,
              isClosable: true,
            });
          }}
        >
          Copy secret link
        </Button>
      </CardBody>
    </Card>
  );
};
