import unittest
import requests

class ApiTests(unittest.Testcase):
  def test_healthcheck(self):
    result = requests.get("http://api:8080/health")
    self.assertEqual(200, result.status_code)
