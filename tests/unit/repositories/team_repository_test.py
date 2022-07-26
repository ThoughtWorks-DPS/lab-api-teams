import os
import json

import boto3
from mypy_boto3_dynamodb.service_resource import DynamoDBServiceResource, Table
import pytest

from moto import mock_dynamodb
from api.models.team import Team


from api.repositories.team_repository import TeamRepository


class TestTeamRepository:

    @pytest.fixture
    def mock_aws_creds(self):
        """Mocked AWS Credentials for moto."""
        os.environ["AWS_ACCESS_KEY_ID"] = "testing"
        os.environ['AWS_DEFAULT_REGION'] = "us-east-2"
        os.environ["AWS_SECRET_ACCESS_KEY"] = "testing"
        os.environ["AWS_SECURITY_TOKEN"] = "testing"
        os.environ["AWS_SESSION_TOKEN"] = "testing"

    @pytest.fixture
    def mock_dynamodb(self, mock_aws_creds):
        with mock_dynamodb():
            conn: DynamoDBServiceResource = boto3.resource('dynamodb')
            yield conn

    @pytest.fixture
    def mock_dynamo_table(self, mock_dynamodb):

        with open('dynamodb_table_schema.json') as schema_file:
            dynamodb_table_schema = schema_file.read()

        schema = json.loads(dynamodb_table_schema)
        table: Table = mock_dynamodb.create_table(
            TableName='teams',
            **schema
        )
        yield table
        table.delete()

    def test_put_team_should_persist_to_dynamodb(self, mock_dynamo_table):
        repository = TeamRepository(mock_dynamo_table)
        test_team = Team(name="dps1")

        repository.put(test_team)
        found_team = mock_dynamo_table.get_item(Key={'name': test_team.name})['Item']['name']

        assert found_team is not None
        assert test_team.name == found_team

    def test_get_team_should_return_team_by_name(self, mock_dynamo_table):
        repository = TeamRepository(mock_dynamo_table)
        test_team = Team(name="dps1")
        mock_dynamo_table.put_item(Item=test_team.dict())
        found_team = repository.get(test_team.name)
        assert found_team is not None
        assert found_team.name == test_team.name

    def test_delete_team_should_return_none_when_already_deleted(self, mock_dynamo_table):
        repository = TeamRepository(mock_dynamo_table)
        test_team = Team(name="something1")
        mock_dynamo_table.put_item(Item=test_team.dict())
        d_one = repository.delete(test_team.name)
        deleted_team = repository.delete(test_team.name)
        assert d_one is True
        assert deleted_team is None 

    def test_delete_team_should_remove_from_dynamodb(self, mock_dynamo_table):
        repository =  TeamRepository(mock_dynamo_table)
        test_team = Team(name="dps1")
        mock_dynamo_table.put_item(Item=test_team.dict())
        found_team = mock_dynamo_table.get_item(Key={'name': test_team.name})['Item']['name']
        assert found_team == "dps1"

        repository.delete(test_team.name)

        assert repository.get(test_team.name) is None

    def test_get_all_should_return_a_list_of_teams(self, mock_dynamo_table):
        repository = TeamRepository(mock_dynamo_table)
        test_team_1 = Team(name="dps1")
        test_team_2 = Team(name="dps2")
        mock_dynamo_table.put_item(Item=test_team_1.dict())
        mock_dynamo_table.put_item(Item=test_team_2.dict())

        found_teams = repository.get_all()

        assert len(found_teams) == 2
        assert test_team_1 in found_teams
        assert test_team_2 in found_teams

    def test_get_all_should_return_empty_list_if_no_teams(self, mock_dynamo_table):
        repository = TeamRepository(mock_dynamo_table)

        found_teams = repository.get_all()

        assert len(found_teams) == 0
