# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from .. import _utilities

import types

__config__ = pulumi.Config('kafka-connect')


class _ExportableConfig(types.ModuleType):
    @property
    def password(self) -> Optional[str]:
        return __config__.get('password')

    @property
    def url(self) -> Optional[str]:
        """
        The url for the kafka connect cluster
        """
        return __config__.get('url')

    @property
    def user(self) -> Optional[str]:
        return __config__.get('user')

