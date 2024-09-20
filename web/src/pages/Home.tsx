import React, { useEffect } from "react";
import {
  Button,
  createStandaloneToast,
  Select,
  Textarea,
} from "@chakra-ui/react";
import { Card, CardBody } from "@chakra-ui/react";
import { Form, json, Link, useActionData, useNavigate } from "react-router-dom";
import axios from "axios";

export async function action({ request, params }) {
  const { toast } = createStandaloneToast();

  const formData = await request.formData();
  const { secret } = Object.fromEntries(formData);

  if (secret === "") {
    toast({
      title: "Secret cannot be empty",
      status: "error",
      duration: 4000,
      isClosable: true,
    });
    return null;
  }

  const { status, data } = await axios.post(
    "http://localhost/api/v1/secret/new",
    {
      secret,
    }
  );

  if (status != 200) {
    toast({
      title: "Something goes wrong",
      status: "error",
      duration: 6000,
      isClosable: true,
    });
    return null;
  }

  return json({
    data: data.data,
    link: "/secret/new",
  });
}

export const Home = () => {
  const actionData = useActionData();
  const navigate = useNavigate();

  useEffect(() => {
    if (!actionData) {
      return
    }

    const { link, data } = actionData as { link: string, data: string }

    navigate(link, { state: { data }, replace: true})
  }, [actionData])

  return (
    <div>
      <h1 className="uppercase text-8xl">
        <Link to={"/"}>🔥 Burn a Secret</Link>
      </h1>
      <p className="text-base mb-6">A link that is created to be burned</p>
      <Form method="post" className="flex flex-col gap-6">
        <Card variant="elevated">
          <CardBody>
            <Textarea placeholder="Secret content goes here..." name="secret" />
          </CardBody>
        </Card>
        <Card variant="elevated">
          <CardBody className="flex gap-8">
            {/* <InputGroup>
            <InputLeftAddon>Passphrase</InputLeftAddon>
            <Input
              name="passphrase"
              placeholder="A world or phrase that is difficult to guess"
            />
          </InputGroup> */}
            <Select placeholder="" name="" disabled>
              <option value="1">Can be seen 1 time</option>
            </Select>
          </CardBody>
        </Card>
        <Button colorScheme="orange" type="submit">
          Create my secret link
        </Button>
      </Form>
    </div>
  );
};
