import React from "react";
import {
  Button,
  Textarea,
} from "@chakra-ui/react";
import { Card, CardBody } from "@chakra-ui/react";
import { Form, Link } from "react-router-dom";

export const Home = () => {
  return (
    <div>
      <h1 className="uppercase text-8xl">
        <Link to={"/"}>Burn a Secret</Link>
      </h1>
      <p className="text-base">
        with a link that is created to be destroyed
      </p>
      <Form method="post" className="flex flex-col gap-6">
        <Card variant="elevated">
          <CardBody>
            <Textarea placeholder="Secret content goes here..." name="secret" />
          </CardBody>
        </Card>
        {/* <Card variant="elevated">
        <CardBody className="flex gap-8">
          <InputGroup>
            <InputLeftAddon>Passphrase</InputLeftAddon>
            <Input
              name="passphrase"
              placeholder="A world or phrase that is difficult to guess"
            />
          </InputGroup>
          <Select placeholder="Lifetime" name="lifetime">
            <option value="option1">7 days</option>
            <option value="option2">3 days</option>
            <option value="option3">1 day</option>
            <option value="option3">12 hours</option>
            <option value="option3">4 hours</option>
            <option value="option3">1 hour</option>
            <option value="option3">30 minutes</option>
            <option value="option3">15 minutes</option>
          </Select>
        </CardBody>
      </Card> */}
        <Button colorScheme="orange" type="submit">
          Create a secret link
        </Button>
      </Form>
    </div>
  );
};
